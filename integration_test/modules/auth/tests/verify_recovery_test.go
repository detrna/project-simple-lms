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

func TestVerifyRecovery_Success(t *testing.T) {
	suite := testsuite.New()

	existingUser := userfactory.CreateUser(t, suite.DB, suite.Infra.Hasher, "student-1")
	otp, err := userfactory.CreateOTP(t, suite.DB, suite, existingUser)
	require.NoError(t, err)

	requestData := auth.VerifyRecoverySchema{
		Email:       existingUser.Email,
		NewPassword: "new-password",
		OTP:         otp,
	}

	requestBody, err := json.Marshal(&requestData)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/auth/verify-recovery",
		bytes.NewReader(requestBody),
	)

	suite.Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestVerifyRecovery_IncorrectOTP(t *testing.T) {
	suite := testsuite.New()

	existingUser := userfactory.CreateUser(t, suite.DB, suite.Infra.Hasher, "student-1")
	_, err := userfactory.CreateOTP(t, suite.DB, suite, existingUser)
	require.NoError(t, err)

	requestData := auth.VerifyRecoverySchema{
		Email:       existingUser.Email,
		NewPassword: "new-password",
		OTP:         "incorrect-otp",
	}

	requestBody, err := json.Marshal(&requestData)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/auth/verify-recovery",
		bytes.NewReader(requestBody),
	)

	suite.Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
