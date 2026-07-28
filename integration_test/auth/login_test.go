package auth_integration_test

import (
	"bytes"
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

func TestLogin_Success(t *testing.T) {
	router, factory := test_suite.SetupSuite()

	existingUser := factory.CreateUser(t, "Student1")

	reqData := auth.LoginSchema{Email: existingUser.Email, Password: "password123"}

	reqBody, err := json.Marshal(&reqData)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(reqBody))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	cookies := w.Result().Cookies()

	assert.Equal(t, 1, len(cookies))
	assert.Equal(t, "refresh_token", cookies[0].Name)

	var response shared.ResponseSuccess[auth.TokenResponse]
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.IsType(t, auth.TokenResponse{}, *response.Data)
	assert.NotNil(t, response.Data.AccessToken)
}

func TestLogin_NonexistentEmail(t *testing.T) {
	router, _ := test_suite.SetupSuite()

	reqData := auth.LoginSchema{Email: "incorrect-email", Password: "password123"}

	reqBody, err := json.Marshal(&reqData)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(reqBody))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	cookies := w.Result().Cookies()

	assert.Equal(t, 0, len(cookies))

	var response shared.ResponseError
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, shared.ErrCredentialsIncorrect.Error(), response.Error)
}

func TestLogin_IncorrectPassword(t *testing.T) {
	router, factory := test_suite.SetupSuite()

	existingUser := factory.CreateUser(t, "Student1")

	reqData := auth.LoginSchema{Email: existingUser.Email, Password: "incorrect-password"}

	reqBody, err := json.Marshal(&reqData)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(reqBody))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	cookies := w.Result().Cookies()

	assert.Equal(t, 0, len(cookies))

	var response shared.ResponseError
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, shared.ErrCredentialsIncorrect.Error(), response.Error)
}
