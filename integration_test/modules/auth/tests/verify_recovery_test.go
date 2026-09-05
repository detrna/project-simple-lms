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

func TestVerifyRecovery_Success(t *testing.T) {
	ts := suite.New()

	existingUser := userfactory.CreateUser(t, ts.DB, ts.Infra.Hasher, "student-1")
	otp, err := userfactory.CreateOTP(t, ts.DB, ts, existingUser)
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

	ts.Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestVerifyRecovery_IncorrectOTP(t *testing.T) {
	ts := suite.New()

	existingUser := userfactory.CreateUser(t, ts.DB, ts.Infra.Hasher, "student-1")
	_, err := userfactory.CreateOTP(t, ts.DB, ts, existingUser)
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

	ts.Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
