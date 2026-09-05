package class_test

import (
	"encoding/json"
	"main/integration_test/helper"
	"main/integration_test/modules/class/factory"
	suite "main/integration_test/suite"
	"main/internal/modules/class/domain"
	"main/internal/modules/class/dto"
	"main/internal/shared"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {
	ts := suite.New()
	existingData := factory.CreateClass(t, ts.DB, "Class-A")

	updatedData := existingData
	updatedData.Name = "class-B"

	tests := []suite.IntegrationTest[dto.UpdateClassRequest]{
		{
			Name: "success",
			Data: dto.UpdateClassRequest{
				ID:   existingData.ID,
				Name: &updatedData.Name,
			},
			ExpectedStatusCode: http.StatusOK,
			ExpectedResponse: shared.ResponseSuccess[domain.Class]{
				Data: existingData,
			},
		},
		{
			Name: "class not found",
			Data: dto.UpdateClassRequest{
				ID:   uuid.New(),
				Name: &updatedData.Name,
			},
			ExpectedStatusCode: http.StatusNotFound,
			ExpectedResponse:   shared.ErrRecordNotFound,
		},
		{
			Name: "systemID taken not found",
			Data: dto.UpdateClassRequest{
				ID:       existingData.ID,
				SystemID: &existingData.SystemID,
				Name:     &updatedData.Name,
			},
			ExpectedStatusCode: http.StatusConflict,
			ExpectedResponse:   shared.ErrSystemIDTaken,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			req := httptest.NewRequest(
				http.MethodPatch,
				"/api/v1/classes",
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
