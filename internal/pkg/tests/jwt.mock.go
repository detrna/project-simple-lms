package mocks

import (
	"main/internal/domain"

	"github.com/stretchr/testify/mock"
)

type MockJWTProvider struct {
	mock.Mock
}

func (m *MockJWTProvider) GenerateAccessToken(
	data *domain.User,
) (*domain.JWT, error) {
	args := m.Called(data)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.JWT), args.Error(1)
}

func (m *MockJWTProvider) GenerateRefreshToken(
	data *domain.User,
) (*domain.JWT, error) {
	args := m.Called(data)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.JWT), args.Error(1)
}

func (m *MockJWTProvider) ParseAccessToken(
	tokenString string,
) (*domain.JWTPayload, error) {
	args := m.Called(tokenString)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.JWTPayload), args.Error(1)
}

func (m *MockJWTProvider) ParseRefreshToken(
	tokenString string,
) (*domain.JWTPayload, error) {
	args := m.Called(tokenString)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.JWTPayload), args.Error(1)
}
