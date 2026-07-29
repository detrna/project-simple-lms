package auth_integration_test

import (
	"bytes"
	"encoding/json"
	test_suite "main/integration_test"
	"main/internal/modules/auth"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRecover_Success(t *testing.T) {
	router, factory := test_suite.SetupSuite()

	existingUser := factory.CreateUser(t, "Student1")

	reqData := auth.RecoverSchema{
		Email: existingUser.Email,
	}

	reqBody, err := json.Marshal(&reqData)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/recover", bytes.NewBuffer(reqBody))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestRecover_NonexistentUser(t *testing.T) {
	router, factory := test_suite.SetupSuite()

	_ = factory.CreateUser(t, "Student1")

	reqData := auth.RecoverSchema{
		Email: "nonexistent-email",
	}

	reqBody, err := json.Marshal(&reqData)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/recover", bytes.NewBuffer(reqBody))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
