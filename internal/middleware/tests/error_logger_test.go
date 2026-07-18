package middleware_test

import (
	"errors"
	"main/internal/middleware"
	mocks "main/internal/pkg/tests"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestErrorLogger_NoError(t *testing.T) {
	mockLogger := new(mocks.MockLogger)

	router := gin.New()

	router.Use(middleware.ErrorLogger(mockLogger))

	router.GET("/users", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	w := httptest.NewRecorder()

	req := httptest.NewRequest(
		http.MethodGet,
		"/users",
		nil,
	)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	mockLogger.AssertNotCalled(
		t,
		"ErrorLog",
	)
}

func TestErrorLogger_WithError(t *testing.T) {
	mockLogger := new(mocks.MockLogger)

	expectedErr := errors.New("something went wrong")

	mockLogger.
		On(
			"ErrorLog",
			http.MethodGet,
			"/users",
			http.StatusInternalServerError,
			expectedErr,
		).
		Once()

	router := gin.New()

	router.Use(middleware.ErrorLogger(mockLogger))

	router.GET("/users", func(c *gin.Context) {
		c.Error(expectedErr)
		c.Status(http.StatusInternalServerError)
	})

	w := httptest.NewRecorder()

	req := httptest.NewRequest(
		http.MethodGet,
		"/users",
		nil,
	)

	router.ServeHTTP(w, req)

	assert.Equal(
		t,
		http.StatusInternalServerError,
		w.Code,
	)

	mockLogger.AssertExpectations(t)
}
