package auth_integration_test

import (
	"encoding/json"
	test_suite "main/integration_test"
	"main/internal/modules/auth"
	"main/internal/shared"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRefresh_Success(t *testing.T) {
	router, factory := test_suite.SetupSuite()

	existingUser := factory.CreateUser(t, "Student1")
	existingJWT := factory.CreateJWT(t, existingUser)
	accessToken, err := factory.Infra.TokenService.GenerateRefreshToken(existingUser)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/refresh", nil)

	req.AddCookie(&http.Cookie{
		Name:  "refresh_token",
		Value: existingJWT.Value,
	})

	req.Header.Set("Authorization", "Bearer "+accessToken.Value)

	router.ServeHTTP(w, req)

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
	router, factory := test_suite.SetupSuite()

	_ = factory.CreateUser(t, "Student1")

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/refresh", nil)

	req.AddCookie(&http.Cookie{
		Name:  "refresh_token",
		Value: "revoked-token",
	})

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var response shared.ResponseError
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, shared.ErrUnauthorized.Error(), response.Error)
}
