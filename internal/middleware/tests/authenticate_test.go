package middleware_test

import (
	"main/internal/domain"
	"main/internal/middleware"
	mocks "main/internal/pkg/tests"
	"main/internal/shared"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticate_MissingAuthorizationHeader(t *testing.T) {
	c, w := shared.SetupTestContext(
		http.MethodGet,
		"/",
		"",
		nil,
	)

	handler := middleware.Authenticate(nil)

	handler(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.True(t, c.IsAborted())
}

func TestAuthenticate_InvalidToken(t *testing.T) {
	header := http.Header{
		"Authorization": []string{"Invalid"},
	}

	c, w := shared.SetupTestContext(
		http.MethodGet,
		"/",
		"",
		header,
	)

	handler := middleware.Authenticate(nil)

	handler(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.True(t, c.IsAborted())
}

func TestAuthenticate_Success(t *testing.T) {
	mockJWT := new(mocks.MockJWTProvider)

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

	router.Use(middleware.Authenticate(mockJWT))

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
