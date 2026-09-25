package class_test

import (
	"fmt"
	"main/integration_test/helper"
	"main/integration_test/modules/auth/factory"
	classfactory "main/integration_test/modules/class/factory"
	suite "main/integration_test/suite"
	"main/internal/modules/class/domain"
	"main/internal/shared"
	"main/internal/shared/pagination"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll_Success(t *testing.T) {
	ts := suite.New()
	existingData := []domain.Class{*classfactory.CreateClass(t, ts.DB, "Class-A")}
	jwt := factory.CreateAdminJWT(t, ts.DB, ts.Infra.TokenService, ts.Infra.Hasher)

	page := 1
	limit := 10
	tests := []suite.IntegrationTest[any]{
		{
			Name:               "success",
			Data:               nil,
			ExpectedStatusCode: http.StatusOK,
			ExpectedResponse: shared.PaginatedResponseSuccess[domain.Class]{
				Data: &existingData,
				Pagination: pagination.GetPaginationResponse(pagination.Pagination{
					Page:  page,
					Limit: limit,
				}, len(existingData)),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			req := httptest.NewRequest(
				http.MethodGet,
				fmt.Sprintf("/api/v1/classes?page=%d&limit=%d", page, limit),
				nil,
			)

			req.Header.Set("Authorization", "Bearer "+jwt.Value)

			w := httptest.NewRecorder()

			ts.Router.ServeHTTP(w, req)
			assert.Equal(t, test.ExpectedStatusCode, w.Code)

			response := helper.ParseJSON(t, w, test.ExpectedResponse)
			assert.Equal(t, test.ExpectedResponse, response)
		})
	}
}
