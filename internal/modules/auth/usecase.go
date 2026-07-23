package auth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"main/internal/config"
	"main/internal/domain"
	"main/internal/modules/user"
	"main/internal/pkg"
	"main/internal/shared"
	"math/big"
	"time"
)

type UseCasePackages struct {
	Bcrypt        pkg.BcryptHasher
	Mailer        pkg.ResendClient
	TokenProvider pkg.JWTProvider
	Redis         pkg.RedisClient
	Logger        pkg.Logger
}

type UseCase struct {
	repo       IRepository
	userRepo   user.IRepository
	packages   UseCasePackages
	mailConfig *config.MailConfig
}

type IUseCase interface {
	Login(ctx context.Context, data *LoginSchema) (*Tokens, error)
	Recover(ctx context.Context, data *RecoverSchema) error
	VerifyRecovery(ctx context.Context, data *VerifyRecoverSchema) error
	Logout(ctx context.Context, refreshToken string) error
	Refresh(ctx context.Context, refreshToken string) (*Tokens, error)
}

func NewUseCase(repo IRepository, userRepo user.IRepository, pkg *UseCasePackages, mailCfg *config.MailConfig) *UseCase {
	return (&UseCase{repo: repo, userRepo: userRepo, packages: *pkg, mailConfig: mailCfg})
}

func (usecase *UseCase) Login(ctx context.Context, data *LoginSchema) (*Tokens, error) {
	dbAccount, err := usecase.userRepo.FindByEmail(ctx, data.Email)

	if err != nil {
		return nil, err
	}

	err = usecase.packages.Bcrypt.Compare(dbAccount.Password, data.Password)

	if err != nil {
		return nil, shared.ErrCredentialsIncorrect
	}

	accessToken, err := usecase.packages.TokenProvider.GenerateAccessToken(dbAccount)
	refreshToken, err := usecase.packages.TokenProvider.GenerateRefreshToken(dbAccount)

	if err != nil {
		return nil, err
	}

	sum := sha256.Sum256([]byte(refreshToken.Value))
	hashedToken := hex.EncodeToString(sum[:])

	_, err = usecase.repo.CreateJWT(ctx, &accessToken.Payload, hashedToken)

	if err != nil {
		return nil, err
	}

	return &Tokens{AccessToken: accessToken.Value, RefreshToken: refreshToken.Value}, nil
}

func (usecase *UseCase) Logout(ctx context.Context, refreshToken string) error {
	tokenPayload, err := usecase.packages.TokenProvider.ParseRefreshToken(refreshToken)

	if err != nil {
		return shared.ErrUnauthorized
	}

	dbToken, err := usecase.repo.FindJWT(ctx, tokenPayload.JTI)

	if err != nil {
		return err
	}

	err = usecase.packages.Bcrypt.Compare(*dbToken, refreshToken)

	if errors.Is(err, shared.ErrCredentialsIncorrect) {
		return shared.ErrUnauthorized
	}

	if err != nil {
		return err
	}

	return usecase.repo.DeleteJWT(ctx, tokenPayload.JTI)
}

func (usecase *UseCase) Refresh(ctx context.Context, refreshToken string) (*Tokens, error) {
	tokenPayload, err := usecase.packages.TokenProvider.ParseRefreshToken(refreshToken)

	if err != nil {
		return nil, shared.ErrUnauthorized
	}

	dbToken, err := usecase.repo.FindJWT(ctx, tokenPayload.JTI)

	if err != nil {
		return nil, err
	}

	err = usecase.packages.Bcrypt.Compare(*dbToken, refreshToken)

	if errors.Is(err, shared.ErrCredentialsIncorrect) {
		return nil, shared.ErrUnauthorized
	}

	if err != nil {
		return nil, err
	}

	user := domain.User{
		ID:       tokenPayload.UserID,
		Role:     tokenPayload.Role,
		Name:     tokenPayload.Name,
		SystemID: tokenPayload.SystemID,
	}

	newAccessToken, err := usecase.packages.TokenProvider.GenerateAccessToken(&user)
	newRefreshToken, err := usecase.packages.TokenProvider.GenerateRefreshToken(&user)

	if err != nil {
		return nil, err
	}

	sum := sha256.Sum256([]byte(newRefreshToken.Value))
	hashedToken := hex.EncodeToString(sum[:])

	_, err = usecase.repo.CreateJWT(ctx, &newRefreshToken.Payload, hashedToken)

	if err != nil {
		return nil, err
	}

	err = usecase.repo.DeleteJWT(ctx, tokenPayload.JTI)

	if err != nil {
		return nil, err
	}

	return &Tokens{AccessToken: newAccessToken.Value, RefreshToken: newRefreshToken.Value}, nil
}

func (usecase UseCase) Recover(ctx context.Context, data *RecoverSchema) error {
	dbAccount, err := usecase.userRepo.FindByEmail(ctx, data.Email)

	if err != nil {
		return err
	}

	otp, _ := rand.Int(rand.Reader, big.NewInt(1000000))

	err = usecase.packages.Mailer.SendRecoveryOTP(ctx, dbAccount, otp.String())

	usecase.packages.Redis.Set(
		ctx,
		"otp:"+dbAccount.Email,
		otp.String(),
		time.Duration(usecase.mailConfig.OTPExpiryMin)*time.Minute,
	)

	if err != nil {
		return err
	}

	return nil
}

func (usecase UseCase) VerifyRecovery(ctx context.Context, data *VerifyRecoverSchema) error {
	existingAccount, err := usecase.userRepo.FindByEmail(ctx, data.Email)

	if err != nil {
		return err
	}

	code, err := usecase.packages.Redis.Get(ctx, data.Email)

	if err != nil {
		return err
	}

	if code != data.OTP {
		return shared.ErrIncorrectOTP
	}

	existingAccount.Password = data.NewPassword

	_, err = usecase.userRepo.Update(ctx, existingAccount)

	return nil
}
