package middleware_test

import (
	"encoding/json"
	"fmt"
	"main/internal/middleware"
	"main/internal/shared/pagination"
	"main/internal/testutil/logger"
	"main/internal/testutil/testcases"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPagination(t *testing.T) {
	const defaultlimit, maxLimit, defaultPage int = 10, 50, 1
	const defaultOffset int = (defaultPage - 1) * defaultlimit

	tests := []testcases.ControllerTest[pagination.Pagination]{
		{
			Name: "max limit is exceeded",
			Data: pagination.Pagination{
				Page:  defaultPage,
				Limit: 100,
			},
			ExpectedStatusCode: http.StatusOK,
			ExpectedResponse: pagination.Pagination{
				Page:   defaultPage,
				Limit:  maxLimit,
				Offset: defaultOffset,
			},
		},
		{
			Name: "max limit is below zero",
			Data: pagination.Pagination{
				Page:  defaultOffset,
				Limit: -1,
			},
			ExpectedStatusCode: http.StatusOK,
			ExpectedResponse: pagination.Pagination{
				Page:   defaultPage,
				Limit:  defaultlimit,
				Offset: defaultOffset,
			},
		},
		{
			Name: "page is lower or lesser zero",
			Data: pagination.Pagination{
				Page:  0,
				Limit: defaultlimit,
			},
			ExpectedStatusCode: http.StatusOK,
			ExpectedResponse: pagination.Pagination{
				Page:   defaultPage,
				Limit:  defaultlimit,
				Offset: defaultOffset,
			},
		},
		{
			Name:               "pagination not requested",
			ExpectedStatusCode: http.StatusOK,
			ExpectedResponse: pagination.Pagination{
				Page:   defaultPage,
				Limit:  defaultlimit,
				Offset: defaultOffset,
			},
		},
		{
			Name: "success",
			Data: pagination.Pagination{
				Page:  5,
				Limit: 15,
			},
			ExpectedStatusCode: http.StatusOK,
			ExpectedResponse: pagination.Pagination{
				Page:   5,
				Limit:  15,
				Offset: (5 - 1) * 15,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			router := gin.New()
			w := httptest.NewRecorder()
			req := httptest.NewRequest(
				http.MethodGet,
				fmt.Sprintf("/?page=%d&limit=%d", test.Data.Page, test.Data.Limit),
				nil,
			)

			router.GET("/", middleware.HandlePagination(defaultlimit, maxLimit, logger.NewMockLogger(t)), func(ctx *gin.Context) {
				paging := ctx.MustGet("pagination")
				ctx.JSON(200, paging)
			})

			router.ServeHTTP(w, req)
			assert.Equal(t, test.ExpectedStatusCode, w.Code)

			var res pagination.Pagination
			err := json.Unmarshal(w.Body.Bytes(), &res)
			require.NoError(t, err)

			assert.Equal(t, test.ExpectedResponse, res)
		})
	}
}
