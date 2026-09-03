package auth_integration_test

import (
	"bytes"
	"encoding/json"
	testsuite "main/integration_test"
	userfactory "main/integration_test/modules/user/factory"
	"main/internal/modules/auth"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRecover_Success(t *testing.T) {
	suite := testsuite.New()

	existingUser := userfactory.CreateUser(t, suite.DB, suite.Infra.Hasher, "Student1")

	reqData := auth.RecoverSchema{
		Email: existingUser.Email,
	}

	reqBody, err := json.Marshal(&reqData)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/recover", bytes.NewBuffer(reqBody))

	suite.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestRecover_NonexistentUser(t *testing.T) {
	suite := testsuite.New()

	_ = userfactory.CreateUser(t, suite.DB, suite.Infra.Hasher, "Student1")

	reqData := auth.RecoverSchema{
		Email: "nonexistent-email",
	}

	reqBody, err := json.Marshal(&reqData)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/recover", bytes.NewBuffer(reqBody))

	suite.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
