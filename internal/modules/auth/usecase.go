package auth

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"main/internal/modules/user"
	"main/internal/shared"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UseCase struct {
	repo     IRepository
	userRepo user.IRepository
	logger   shared.Logger
}

type IUseCase interface {
	Login(ctx context.Context, data LoginSchema) (*Tokens, error)
	// Recover(ctx context.Context, data RecoverSchema) (*Tokens, error)
	// Logout(ctx context.Context, id uuid.UUID) error
	// Refresh(ctx context.Context, token string) (*Tokens, error)
}

func NewUseCase(repo IRepository, userRepo user.IRepository, logger shared.Logger) *UseCase {
	return (&UseCase{repo: repo, userRepo: userRepo, logger: logger})
}

func (usecase UseCase) Login(ctx context.Context, data LoginSchema) (*Tokens, error) {
	dbAccount, err := usecase.userRepo.FindByEmail(ctx, data.Email)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbAccount.Password), []byte(data.Password))

	if err != nil {
		return nil, shared.ErrCredentialsIncorrect
	}

	jwtPayload := shared.JWTPayload{
		UserId:   dbAccount.ID.String(),
		SystemId: dbAccount.SystemID,
		Name:     dbAccount.Name,
		Role:     dbAccount.Role,
	}

	accessToken, err := shared.GenerateAccessToken(jwtPayload)
	refreshToken, err := shared.GenerateRefreshToken(jwtPayload)

	if err != nil {
		return nil, err
	}

	sum := sha256.Sum256([]byte(refreshToken))
	hashedToken := hex.EncodeToString(sum[:])

	dbJWT := JWT{
		ID:        uuid.New(),
		UserID:    dbAccount.ID,
		Value:     string(hashedToken),
		CreatedAt: time.Now(),
	}

	result, err := usecase.repo.CreateJWT(ctx, dbJWT)

	return &Tokens{AccessToken: accessToken, RefreshToken: *result}, nil
}
