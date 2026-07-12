package auth

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"main/internal/domain"
	"main/internal/modules/user"
	"main/internal/pkg"
	"main/internal/shared"

	"github.com/google/uuid"
)

type UseCase struct {
	repo     IRepository
	userRepo user.IRepository
	infra    pkg.Packages
}

type IUseCase interface {
	Login(ctx context.Context, data LoginSchema) (*Tokens, error)
	Recover(ctx context.Context, data RecoverSchema) error
	VerifyRecovery(ctx context.Context, data VerifyRecoverSchema) (*domain.User, error)
	Logout(ctx context.Context, id uuid.UUID) error
	Refresh(ctx context.Context, JWTPayload domain.JWTPayload) (*Tokens, error)
}

func NewUseCase(repo IRepository, userRepo user.IRepository, infra pkg.Packages) *UseCase {
	return (&UseCase{repo: repo, userRepo: userRepo, infra: infra})
}

func (usecase *UseCase) Login(ctx context.Context, data LoginSchema) (*Tokens, error) {
	dbAccount, err := usecase.userRepo.FindByEmail(ctx, data.Email)

	if err != nil {
		return nil, err
	}

	err = usecase.infra.BcryptHasher.CompareHashAndPassword(dbAccount.Password, data.Password)

	if err != nil {
		return nil, shared.ErrCredentialsIncorrect
	}

	JWTPayload := domain.JWTPayload{
		JTI:      uuid.New(),
		UserID:   dbAccount.ID,
		SystemID: dbAccount.SystemID,
		Name:     dbAccount.Name,
		Role:     dbAccount.Role,
	}

	accessToken, err := usecase.infra.JWTProvider.GenerateAccessToken(JWTPayload)
	refreshToken, err := usecase.infra.JWTProvider.GenerateRefreshToken(JWTPayload)

	if err != nil {
		return nil, err
	}

	sum := sha256.Sum256([]byte(refreshToken))
	hashedToken := hex.EncodeToString(sum[:])

	result, err := usecase.repo.CreateJWT(ctx, JWTPayload, hashedToken)

	if err != nil {
		return nil, err
	}

	return &Tokens{AccessToken: accessToken, RefreshToken: *result}, nil
}

func (usecase *UseCase) Logout(ctx context.Context, ID uuid.UUID) error {
	err := usecase.repo.DeleteJWT(ctx, ID)

	if err != nil {
		return err
	}

	return nil
}

func (usecase *UseCase) Refresh(ctx context.Context, JWTPayload domain.JWTPayload) (*Tokens, error) {
	err := usecase.repo.CheckJWT(ctx, JWTPayload.JTI)

	if err != nil {
		return nil, err
	}

	accessToken, err := usecase.infra.JWTProvider.GenerateAccessToken(JWTPayload)
	refreshToken, err := usecase.infra.JWTProvider.GenerateRefreshToken(JWTPayload)

	if err != nil {
		return nil, err
	}

	sum := sha256.Sum256([]byte(refreshToken))
	hashedToken := hex.EncodeToString(sum[:])

	result, err := usecase.repo.CreateJWT(ctx, JWTPayload, hashedToken)

	if err != nil {
		return nil, err
	}

	return &Tokens{AccessToken: accessToken, RefreshToken: *result}, nil
}

func (usecase UseCase) Recover(ctx context.Context, data RecoverSchema) error {
	dbAccount, err := usecase.userRepo.FindByEmail(ctx, data.Email)

	if err != nil {
		return err
	}

	err = usecase.infra.ResendClient.SendRecoveryOTP(ctx, *dbAccount)

	if err != nil {
		return err
	}

	return nil
}

func (usecase UseCase) VerifyRecovery(ctx context.Context, data VerifyRecoverSchema) (*domain.User, error) {
	code, err := usecase.infra.RedisClient.Get(ctx, data.Email)

	if err != nil {
		return nil, err
	}

	if code != data.OTP {
		return nil, shared.ErrIncorrectOTP
	}

	dbAccount, _ := usecase.userRepo.FindByEmail(ctx, data.Email)

	dbAccount.Password = data.Password

	newAccount, err := usecase.userRepo.Update(ctx, *dbAccount)

	return newAccount, nil
}
