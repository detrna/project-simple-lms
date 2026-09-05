package auth_integration_test

import (
	"bytes"
	"encoding/json"
	userfactory "main/integration_test/modules/user/factory"
	suite "main/integration_test/suite"
	"main/internal/modules/auth"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRecover_Success(t *testing.T) {
	ts := suite.New()

	existingUser := userfactory.CreateUser(t, ts.DB, ts.Infra.Hasher, "Student1")

	reqData := auth.RecoverSchema{
		Email: existingUser.Email,
	}

	reqBody, err := json.Marshal(&reqData)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/recover", bytes.NewBuffer(reqBody))

	ts.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestRecover_NonexistentUser(t *testing.T) {
	ts := suite.New()

	_ = userfactory.CreateUser(t, ts.DB, ts.Infra.Hasher, "Student1")

	reqData := auth.RecoverSchema{
		Email: "nonexistent-email",
	}

	reqBody, err := json.Marshal(&reqData)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/recover", bytes.NewBuffer(reqBody))

	ts.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
