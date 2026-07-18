package shared

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupTestContext(
	method string,
	target string,
	body string,
	headers http.Header,
) (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest(
		method,
		target,
		strings.NewReader(body),
	)

	c.Request.Header = headers

	return c, w
}
