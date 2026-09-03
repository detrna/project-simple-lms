package class_test

import (
	"encoding/json"
	testsuite "main/integration_test"
	"main/integration_test/modules/class/factory"
	"main/internal/modules/class/domain"
	"main/internal/shared"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetBySystemID(t *testing.T) {
	suite := testsuite.New()
	existingData := factory.CreateClass(t, suite.DB, "Class-A")

	tests := []testsuite.IntegrationTest[string]{
		{
			Name:               "success",
			Data:               existingData.SystemID,
			ExpectedStatusCode: http.StatusOK,
			ExpectedResponse: shared.ResponseSuccess[domain.Class]{
				Data: existingData,
			},
		},
		{
			Name:               "class not found",
			Data:               "nonexistent-class-systemID",
			ExpectedStatusCode: http.StatusNotFound,
			ExpectedResponse:   shared.ErrRecordNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			req := httptest.NewRequest(
				http.MethodGet,
				"/api/v1/classes/system"+test.Data,
				nil,
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
