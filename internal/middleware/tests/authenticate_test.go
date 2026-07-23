package middleware_test

import (
	"main/internal/domain"
	"main/internal/middleware"
	pkg_mocks "main/internal/pkg/mocks"
	"main/internal/shared"
	shared_testing "main/internal/shared/testing_helper"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticate_MissingAuthorizationHeader(t *testing.T) {
	mockLogger := shared_testing.NewMockLogger(t)

	c, w := shared.SetupTestContext(
		http.MethodGet,
		"/",
		"",
		nil,
	)

	handler := middleware.Authenticate(nil, mockLogger)

	handler(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.True(t, c.IsAborted())
}

func TestAuthenticate_InvalidToken(t *testing.T) {
	mockLogger := shared_testing.NewMockLogger(t)

	header := http.Header{
		"Authorization": []string{"Invalid"},
	}

	c, w := shared.SetupTestContext(
		http.MethodGet,
		"/",
		"",
		header,
	)

	handler := middleware.Authenticate(nil, mockLogger)

	handler(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.True(t, c.IsAborted())
}

func TestAuthenticate_Success(t *testing.T) {
	mockJWT := new(pkg_mocks.MockJWTProvider)
	mockLogger := shared_testing.NewMockLogger(t)

	expectedPayload := &domain.JWTPayload{
		UserID:   uuid.New(),
		SystemID: "test",
		Name:     "test",
		Role:     "test",
	}

	mockJWT.
		On("ParseAccessToken", "valid-token").
		Return(expectedPayload, nil)

	router := gin.New()

	router.Use(middleware.Authenticate(mockJWT, mockLogger))

	router.GET("/", func(c *gin.Context) {
		user, exists := c.Get("user")

		assert.True(t, exists)
		assert.Equal(t, expectedPayload, user)

		c.Status(http.StatusOK)
	})

	w := httptest.NewRecorder()

	req := httptest.NewRequest(
		http.MethodGet,
		"/",
		nil,
	)

	req.Header.Set(
		"Authorization",
		"Bearer valid-token",
	)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	mockJWT.AssertExpectations(t)
}
