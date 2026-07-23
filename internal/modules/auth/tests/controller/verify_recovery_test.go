package auth_controller_test

import (
	"bytes"
	"context"
	"encoding/json"
	"main/internal/modules/auth"
	auth_mocks "main/internal/modules/auth/mocks"
	user_factory "main/internal/modules/user/tests"
	"main/internal/shared"
	shared_testing "main/internal/shared/testing_helper"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestVerifyRecovery_Success(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	existingUser := user_factory.NewUser(id)

	mockUsecase := auth_mocks.NewMockIUseCase(t)
	mockLogger := shared_testing.NewMockLogger(t)

	mockUsecase.EXPECT().VerifyRecovery(ctx, mock.AnythingOfType("*auth.VerifyRecoverySchema")).Return(nil)

	requestData := auth.VerifyRecoverySchema{
		Email:       existingUser.Email,
		NewPassword: "new-password",
		OTP:         "correct-otp-code",
	}

	requestBody, err := json.Marshal(requestData)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/verify-recovery",
		bytes.NewReader(requestBody),
	)

	req.Header.Set("Content-Type", "application/json")

	router := gin.New()
	c := auth.NewController(mockUsecase, mockLogger, false)

	router.POST("/verify-recovery", func(ctx *gin.Context) {
		c.VerifyRecovery(ctx)
	})

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestVerifyRecovery_IncorrectOTP(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	existingUser := user_factory.NewUser(id)

	mockUsecase := auth_mocks.NewMockIUseCase(t)
	mockLogger := shared_testing.NewMockLogger(t)

	mockUsecase.
		EXPECT().
		VerifyRecovery(ctx, mock.AnythingOfType("*auth.VerifyRecoverySchema")).
		Return(shared.ErrIncorrectOTP)

	requestData := auth.VerifyRecoverySchema{
		Email:       existingUser.Email,
		NewPassword: "new-password",
		OTP:         "incorrect-otp-code",
	}

	requestBody, err := json.Marshal(requestData)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/verify-recovery",
		bytes.NewReader(requestBody),
	)

	req.Header.Set("Content-Type", "application/json")

	router := gin.New()
	c := auth.NewController(mockUsecase, mockLogger, false)

	router.POST("/verify-recovery", func(ctx *gin.Context) {
		c.VerifyRecovery(ctx)
	})

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var response shared.ResponseError
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, shared.ErrIncorrectOTP.Error(), response.Error)
}

func TestVerifyRecovery_OTPNotFound(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	existingUser := user_factory.NewUser(id)

	mockUsecase := auth_mocks.NewMockIUseCase(t)
	mockLogger := shared_testing.NewMockLogger(t)

	mockUsecase.
		EXPECT().
		VerifyRecovery(ctx, mock.AnythingOfType("*auth.VerifyRecoverySchema")).
		Return(shared.ErrRedisRecordNotFound)

	requestData := auth.VerifyRecoverySchema{
		Email:       existingUser.Email,
		NewPassword: "new-password",
		OTP:         "non-existent-otp-code",
	}

	requestBody, err := json.Marshal(requestData)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/verify-recovery",
		bytes.NewReader(requestBody),
	)

	req.Header.Set("Content-Type", "application/json")

	router := gin.New()
	c := auth.NewController(mockUsecase, mockLogger, false)

	router.POST("/verify-recovery", func(ctx *gin.Context) {
		c.VerifyRecovery(ctx)
	})

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var response shared.ResponseError
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, shared.ErrRedisRecordNotFound.Error(), response.Error)

}
