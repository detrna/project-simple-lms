package auth_integration_test

import (
	testsuite "main/integration_test"
	authfactory "main/integration_test/modules/auth/factory"
	userfactory "main/integration_test/modules/user/factory"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLogout_Success(t *testing.T) {
	suite := testsuite.New()

	existingUser := userfactory.CreateUser(t, suite.DB, suite.Infra.Hasher, "Student1")
	existingJWT := authfactory.CreateJWT(t, suite.DB, suite.Infra.TokenService, existingUser)

	accessToken, err := suite.Infra.TokenService.GenerateRefreshToken(existingUser)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/auth/logout", nil)

	req.AddCookie(&http.Cookie{
		Name:  "refresh_token",
		Value: existingJWT.Value,
	})

	req.Header.Set("Authorization", "Bearer "+accessToken.Value)

	suite.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
