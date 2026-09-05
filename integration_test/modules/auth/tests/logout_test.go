package auth_integration_test

import (
	authfactory "main/integration_test/modules/auth/factory"
	userfactory "main/integration_test/modules/user/factory"
	suite "main/integration_test/suite"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLogout_Success(t *testing.T) {
	ts := suite.New()

	existingUser := userfactory.CreateUser(t, ts.DB, ts.Infra.Hasher, "Student1")
	existingJWT := authfactory.CreateJWT(t, ts.DB, ts.Infra.TokenService, existingUser)

	accessToken, err := ts.Infra.TokenService.GenerateRefreshToken(existingUser)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/auth/logout", nil)

	req.AddCookie(&http.Cookie{
		Name:  "refresh_token",
		Value: existingJWT.Value,
	})

	req.Header.Set("Authorization", "Bearer "+accessToken.Value)

	ts.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
