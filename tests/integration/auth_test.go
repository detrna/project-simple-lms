package integration_test

import (
	"bytes"
	"encoding/json"
	"main/internal/modules/auth"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {
	targettedUser := Factory.CreateUser(t, "Student1")

	reqData := auth.LoginSchema{Email: targettedUser.Email, Password: "password123"}

	reqBody, err := json.Marshal(reqData)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/auth/login", bytes.NewBuffer(reqBody))

	w := httptest.NewRecorder()

	Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
