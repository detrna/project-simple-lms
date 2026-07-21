package middleware_test

import (
	"main/internal/middleware"
	pkg_mocks "main/internal/pkg/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRequestLogger_FailedRequest(t *testing.T) {
	mockLogger := new(pkg_mocks.MockLogger)

	mockLogger.
		On(
			"RequestLog",
			mock.AnythingOfType("string"),
			http.MethodGet,
			"/tests",
			http.StatusInternalServerError,
			mock.AnythingOfType("time.Time"),
		).
		Once()

	router := gin.New()

	router.Use(middleware.RequestLogger(mockLogger))

	router.GET("/tests", func(c *gin.Context) {
		c.Status(http.StatusInternalServerError)
	})

	w := httptest.NewRecorder()

	req := httptest.NewRequest(
		http.MethodGet,
		"/tests",
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

func TestRequestLogger_SuccessRequest(t *testing.T) {
	mockLogger := new(pkg_mocks.MockLogger)

	mockLogger.
		On(
			"RequestLog",
			mock.AnythingOfType("string"),
			http.MethodGet,
			"/tests",
			http.StatusOK,
			mock.AnythingOfType("time.Time"),
		).
		Once()

	router := gin.New()

	router.Use(middleware.RequestLogger(mockLogger))

	router.GET("/tests", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	w := httptest.NewRecorder()

	req := httptest.NewRequest(
		http.MethodGet,
		"/tests",
		nil,
	)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	mockLogger.AssertExpectations(t)
}
