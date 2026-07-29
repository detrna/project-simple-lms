package auth

import (
	"context"
	"crypto/rand"
	"errors"
	"main/internal/config"
	"main/internal/domain"
	"main/internal/modules/user"
	"main/internal/pkg"
	"main/internal/shared"
	"main/internal/shared/templates"
	"math/big"
	"strconv"
	"time"
)

type UseCasePackages struct {
	Hasher       pkg.Hasher
	Mailer       pkg.Mailer
	TokenService pkg.TokenService
	Redis        pkg.RedisClient
	Logger       pkg.Logger
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
	VerifyRecovery(ctx context.Context, data *VerifyRecoverySchema) error
	Logout(ctx context.Context, refreshToken string) error
	Refresh(ctx context.Context, refreshToken string) (*Tokens, error)
}

func NewUseCase(repo IRepository, userRepo user.IRepository, pkg *UseCasePackages, mailCfg *config.MailConfig) *UseCase {
	return (&UseCase{repo: repo, userRepo: userRepo, packages: *pkg, mailConfig: mailCfg})
}

func (usecase *UseCase) Login(ctx context.Context, data *LoginSchema) (*Tokens, error) {
	dbAccount, err := usecase.userRepo.FindByEmail(ctx, data.Email)

	if dbAccount == nil {
		return nil, shared.ErrCredentialsIncorrect
	}

	if !errors.Is(err, shared.ErrRecordNotFound) && err != nil {
		return nil, err
	}

	err = usecase.packages.Hasher.Compare(dbAccount.Password, data.Password)

	if err != nil {
		return nil, shared.ErrCredentialsIncorrect
	}

	accessToken, err := usecase.packages.TokenService.GenerateAccessToken(dbAccount)
	refreshToken, err := usecase.packages.TokenService.GenerateRefreshToken(dbAccount)

	if err != nil {
		return nil, err
	}

	hashedToken := usecase.packages.TokenService.HashToken(refreshToken.Value)

	_, err = usecase.repo.CreateJWT(ctx, &accessToken.Payload, hashedToken)

	if err != nil {
		return nil, err
	}

	return &Tokens{AccessToken: accessToken.Value, RefreshToken: refreshToken.Value}, nil
}

func (usecase *UseCase) Logout(ctx context.Context, refreshToken string) error {
	tokenPayload, err := usecase.packages.TokenService.ParseRefreshToken(refreshToken)

	if err != nil {
		return shared.ErrUnauthorized
	}

	dbToken, err := usecase.repo.FindJWT(ctx, tokenPayload.JTI)

	if err != nil {
		return err
	}

	if !usecase.packages.TokenService.Compare(*dbToken, refreshToken) {
		return shared.ErrUnauthorized
	}

	return usecase.repo.DeleteJWT(ctx, tokenPayload.JTI)
}

func (usecase *UseCase) Refresh(ctx context.Context, refreshToken string) (*Tokens, error) {
	tokenPayload, err := usecase.packages.TokenService.ParseRefreshToken(refreshToken)

	if err != nil {
		return nil, shared.ErrUnauthorized
	}

	dbToken, err := usecase.repo.FindJWT(ctx, tokenPayload.JTI)

	if err != nil {
		return nil, err
	}

	if !usecase.packages.TokenService.Compare(*dbToken, refreshToken) {
		return nil, shared.ErrUnauthorized
	}

	user := domain.User{
		ID:       tokenPayload.UserID,
		Role:     tokenPayload.Role,
		Name:     tokenPayload.Name,
		SystemID: tokenPayload.SystemID,
	}

	newAccessToken, err := usecase.packages.TokenService.GenerateAccessToken(&user)
	newRefreshToken, err := usecase.packages.TokenService.GenerateRefreshToken(&user)

	if err != nil {
		return nil, err
	}

	hashedToken := usecase.packages.TokenService.HashToken(newRefreshToken.Value)

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

	usecase.packages.Logger.Info(usecase.mailConfig.Host)

	otp, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	htmlBody, err := templates.ResetPassword(templates.ResetPasswordDTO{
		Name:   dbAccount.Name,
		OTP:    otp.String(),
		Expiry: strconv.Itoa(usecase.mailConfig.OTPExpiryMin),
	})

	if err != nil {
		return err
	}

	if err = usecase.packages.Mailer.Send(
		ctx,
		dbAccount.Email,
		"Reset your Password",
		htmlBody,
	); err != nil {
		return err
	}

	if err = usecase.packages.Redis.Set(
		ctx,
		"otp:"+dbAccount.Email,
		otp.String(),
		time.Duration(usecase.mailConfig.OTPExpiryMin)*time.Minute,
	); err != nil {
		usecase.packages.Logger.Info(err.Error())
		return err
	}

	return nil
}

func (usecase UseCase) VerifyRecovery(ctx context.Context, data *VerifyRecoverySchema) error {
	existingAccount, err := usecase.userRepo.FindByEmail(ctx, data.Email)

	if existingAccount == nil {
		return shared.ErrRecordNotFound
	}

	if err != nil {
		return err
	}

	code, err := usecase.packages.Redis.Get(ctx, "otp:"+data.Email)

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
