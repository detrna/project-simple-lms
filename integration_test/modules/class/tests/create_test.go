package class_test

import (
	"encoding/json"
	testsuite "main/integration_test"
	"main/integration_test/helper"
	"main/integration_test/modules/class/factory"
	"main/internal/modules/class/domain"
	"main/internal/modules/class/dto"
	"main/internal/shared"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	suite := testsuite.New()
	sampleData := factory.CreateClass(t, suite.DB, "Class-A")

	tests := []testsuite.IntegrationTest[dto.CreateClassRequest]{
		{
			Name: "When create class and success",
			Data: dto.CreateClassRequest{
				SystemID: sampleData.SystemID,
				Name:     sampleData.Name,
			},
			ExpectedStatusCode: http.StatusCreated,
			ExpectedResponse: shared.ResponseSuccess[domain.Class]{
				Data: sampleData,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			req := httptest.NewRequest(
				http.MethodPost,
				"/api/v1/classes",
				helper.StructToJSON(t, &test.Data),
			)

			w := httptest.NewRecorder()

			suite.Router.ServeHTTP(w, req)
			assert.Equal(t, test.ExpectedStatusCode, w.Code)

			responseType := reflect.TypeOf(test.ExpectedResponse)
			response := reflect.New(responseType).Interface()

			err := json.Unmarshal(w.Body.Bytes(), response)
			require.NoError(t, err)

			assert.Equal(t, test.ExpectedResponse, &response)
		})
	}
}
