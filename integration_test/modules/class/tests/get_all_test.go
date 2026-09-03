package class_test

import (
	"encoding/json"
	"fmt"
	testsuite "main/integration_test"
	"main/integration_test/helper"
	"main/integration_test/modules/class/factory"
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
	suite := testsuite.New()
	existingData := []domain.Class{*factory.CreateClass(t, suite.DB, "Class-A")}

	paginationRequest := pagination.PaginationRequest{
		Page:  0,
		Limit: 10,
	}

	tests := []testsuite.IntegrationTest[pagination.PaginationRequest]{
		{
			Name:               "success",
			Data:               paginationRequest,
			ExpectedStatusCode: http.StatusOK,
			ExpectedResponse: shared.PaginatedResponseSuccess[domain.Class]{
				Data:       &existingData,
				Pagination: pagination.GetPaginationResponse(paginationRequest, len(existingData)),
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

			suite.Router.ServeHTTP(w, req)
			assert.Equal(t, test.ExpectedStatusCode, w.Code)

			responseType := reflect.TypeOf(test.ExpectedResponse)
			response := reflect.New(responseType).Interface()

			err := json.Unmarshal(w.Body.Bytes(), response)
			require.NoError(t, err)

			assert.Equal(t, test.ExpectedResponse, response)
		})
	}
}
