package auth_integration_test

import (
	"encoding/json"
	authfactory "main/integration_test/modules/auth/factory"
	userfactory "main/integration_test/modules/user/factory"
	suite "main/integration_test/suite"
	"main/internal/modules/auth"
	"main/internal/shared"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRefresh_Success(t *testing.T) {
	ts := suite.New()

	existingUser := userfactory.CreateUser(t, ts.DB, ts.Infra.Hasher, "Student1")
	existingJWT := authfactory.CreateJWT(t, ts.DB, ts.Infra.TokenService, existingUser)
	accessToken, err := ts.Infra.TokenService.GenerateRefreshToken(existingUser)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/refresh", nil)

	req.AddCookie(&http.Cookie{
		Name:  "refresh_token",
		Value: existingJWT.Value,
	})

	req.Header.Set("Authorization", "Bearer "+accessToken.Value)

	ts.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	cookies := w.Result().Cookies()

	assert.Equal(t, 1, len(cookies))
	assert.Equal(t, "refresh_token", cookies[0].Name)
	assert.NotEqual(t, existingJWT.Value, cookies[0].Value)

	assert.NotEqual(t, accessToken, w.Header().Get("Authorization"))

	var response shared.ResponseSuccess[auth.TokenResponse]
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.IsType(t, auth.TokenResponse{}, *response.Data)
	assert.NotNil(t, response.Data.AccessToken)
}

func TestRefresh_RevokedToken(t *testing.T) {
	ts := suite.New()

	_ = userfactory.CreateUser(t, ts.DB, ts.Infra.Hasher, "Student1")

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/refresh", nil)

	req.AddCookie(&http.Cookie{
		Name:  "refresh_token",
		Value: "revoked-token",
	})

	ts.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var response shared.ResponseError
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, shared.ErrUnauthorized.Error(), response.Error)
}
