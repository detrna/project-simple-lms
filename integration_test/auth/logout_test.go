package auth_integration_test

import (
	test_suite "main/integration_test"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLogout_Success(t *testing.T) {
	router, factory := test_suite.SetupSuite()

	existingUser := factory.CreateUser(t, "Student1")
	existingJWT := factory.CreateJWT(t, existingUser)

	accessToken, err := factory.Infra.JWTProvider.GenerateRefreshToken(existingUser)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/auth/logout", nil)

	req.AddCookie(&http.Cookie{
		Name:  "refresh_token",
		Value: existingJWT.Value,
	})

	req.Header.Set("Authorization", "Bearer "+accessToken.Value)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
