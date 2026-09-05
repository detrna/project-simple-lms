package class_test

import (
	"encoding/json"
	"fmt"
	"main/integration_test/helper"
	"main/integration_test/modules/class/factory"
	suite "main/integration_test/suite"
	"main/internal/modules/class/domain"
	"main/internal/shared"
	"main/internal/shared/pagination"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAll_Success(t *testing.T) {
	ts := suite.New()
	existingData := []domain.Class{*factory.CreateClass(t, ts.DB, "Class-A")}

	page := 0
	limit := 10
	paginationRequest := pagination.Pagination{
		Page:  page,
		Limit: limit,
	}

	tests := []suite.IntegrationTest[pagination.Pagination]{
		{
			Name:               "success",
			Data:               paginationRequest,
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
				fmt.Sprintf("/api/v1/classes?page=%d&limit=%d,", paginationRequest.Page, paginationRequest.Limit),
				helper.StructToJSON(t, &test.Data),
			)

			w := httptest.NewRecorder()

			ts.Router.ServeHTTP(w, req)
			assert.Equal(t, test.ExpectedStatusCode, w.Code)

			responseType := reflect.TypeOf(test.ExpectedResponse)
			response := reflect.New(responseType).Interface()

			err := json.Unmarshal(w.Body.Bytes(), response)
			require.NoError(t, err)

			assert.Equal(t, test.ExpectedResponse, response)
		})
	}
}
