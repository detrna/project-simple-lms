package auth_integration_test

import (
	"bytes"
	"encoding/json"
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

func TestLogin_Success(t *testing.T) {
	ts := suite.New()

	existingUser := userfactory.CreateUser(t, ts.DB, ts.Infra.Hasher, "Student1")

	reqData := auth.LoginSchema{Email: existingUser.Email, Password: "password123"}

	reqBody, err := json.Marshal(&reqData)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(reqBody))

	ts.Router.ServeHTTP(w, req)

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
	ts := suite.New()

	reqData := auth.LoginSchema{Email: "incorrect-email", Password: "password123"}

	reqBody, err := json.Marshal(&reqData)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(reqBody))

	ts.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	cookies := w.Result().Cookies()

	assert.Equal(t, 0, len(cookies))

	var response shared.ResponseError
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, shared.ErrCredentialsIncorrect.Error(), response.Error)
}

func TestLogin_IncorrectPassword(t *testing.T) {
	ts := suite.New()

	existingUser := userfactory.CreateUser(t, ts.DB, ts.Infra.Hasher, "Student1")

	reqData := auth.LoginSchema{Email: existingUser.Email, Password: "incorrect-password"}

	reqBody, err := json.Marshal(&reqData)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(reqBody))

	ts.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	cookies := w.Result().Cookies()

	assert.Equal(t, 0, len(cookies))

	var response shared.ResponseError
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, shared.ErrCredentialsIncorrect.Error(), response.Error)
}
