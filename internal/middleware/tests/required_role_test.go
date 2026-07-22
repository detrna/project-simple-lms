package middleware_test

import (
	"encoding/json"
	"main/internal/domain"
	"main/internal/middleware"
	"main/internal/shared"
	shared_testing "main/internal/shared/testing_helper"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRequiredRole_Success(t *testing.T) {
	mockLogger := shared_testing.NewMockLogger(t)

	router := gin.New()
	router.GET(
		"/",
		func(ctx *gin.Context) {
			ctx.Set("user", domain.JWTPayload{Role: "admin"})
		},
		middleware.RequiredRole("admin", mockLogger),
		func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, nil)
		},
	)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodGet,
		"/",
		nil,
	)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRequiredRole_RoleMismatch(t *testing.T) {
	mockLogger := shared_testing.NewMockLogger(t)

	router := gin.New()
	router.GET(
		"/",
		func(ctx *gin.Context) {
			ctx.Set("user", domain.JWTPayload{Role: "default"})
		},
		middleware.RequiredRole("admin", mockLogger),
		func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, nil)
		},
	)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodGet,
		"/",
		nil,
	)

	router.ServeHTTP(w, req)

	var response shared.ResponseError
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.Equal(t, shared.ErrForbidden.Error(), response.Error)
}
