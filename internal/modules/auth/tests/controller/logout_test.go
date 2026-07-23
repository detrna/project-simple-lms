package auth_controller_test

import (
	"context"
	"main/internal/modules/auth"
	auth_mocks "main/internal/modules/auth/mocks"
	user_factory "main/internal/modules/user/tests"
	shared_testing "main/internal/shared/testing_helper"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogout_Success(t *testing.T) {
	ctx := context.Background()

	id := uuid.New()
	existingUser := user_factory.NewUser(id)
	jwtPayload := user_factory.NewJWTPayload(existingUser)
	refreshToken := user_factory.NewJWT("refresh-token", jwtPayload)

	mockUsecase := auth_mocks.NewMockIUseCase(t)
	mockLogger := shared_testing.NewMockLogger(t)

	mockUsecase.EXPECT().Logout(ctx, mock.AnythingOfType("string")).Return(nil)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodDelete,
		"/logout",
		nil,
	)

	req.AddCookie(&http.Cookie{
		Name:  "refresh_token",
		Value: refreshToken.Value,
	})

	router := gin.New()
	c := auth.NewController(mockUsecase, mockLogger, false)

	router.DELETE("/logout", func(ctx *gin.Context) {
		c.Logout(ctx)
	})

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
