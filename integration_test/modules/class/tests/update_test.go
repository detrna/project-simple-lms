package class_test

import (
	"main/integration_test/helper"
	authfactory "main/integration_test/modules/auth/factory"
	classfactory "main/integration_test/modules/class/factory"
	suite "main/integration_test/suite"
	"main/internal/modules/class/domain"
	"main/internal/modules/class/dto"
	"main/internal/shared"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	ts := suite.New()
	existingData := classfactory.CreateClass(t, ts.DB, "Class-A")
	otherExistingData := classfactory.CreateClass(t, ts.DB, "Class-B")
	jwt := authfactory.CreateAdminJWT(t, ts.DB, ts.Infra.TokenService, ts.Infra.Hasher)

	updatedData := existingData
	updatedData.Name = "Class-C"

	tests := []suite.IntegrationTest[dto.UpdateClassRequest]{
		{
			Name: "success",
			Data: dto.UpdateClassRequest{
				ID:   existingData.ID,
				Name: &updatedData.Name,
			},
			ExpectedStatusCode: http.StatusOK,
			ExpectedResponse: shared.ResponseSuccess[domain.Class]{
				Data: &domain.Class{
					ID:       existingData.ID,
					SystemID: existingData.SystemID,
					Name:     updatedData.Name,
				},
			},
		},
		{
			Name: "class not found",
			Data: dto.UpdateClassRequest{
				ID:   uuid.New(),
				Name: &updatedData.Name,
			},
			ExpectedStatusCode: http.StatusNotFound,
			ExpectedResponse: shared.ResponseError{
				Error:   domain.ErrClassNotFound.Error(),
				Message: domain.ErrClassNotFound.Message,
			},
		},
		{
			Name: "systemID taken",
			Data: dto.UpdateClassRequest{
				ID:       existingData.ID,
				SystemID: &otherExistingData.SystemID,
				Name:     &updatedData.Name,
			},
			ExpectedStatusCode: http.StatusConflict,
			ExpectedResponse: shared.ResponseError{
				Error:   domain.ErrClassSystemIDTaken.Error(),
				Message: domain.ErrClassSystemIDTaken.Message,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			req := httptest.NewRequest(
				http.MethodPatch,
				"/api/v1/classes/"+(test.Data.ID).String(),
				helper.StructToJSON(t, &test.Data),
			)
			req.Header.Set("Authorization", "Bearer "+jwt.Value)

			w := httptest.NewRecorder()

			ts.Router.ServeHTTP(w, req)
			assert.Equal(t, test.ExpectedStatusCode, w.Code)

			responseRaw := helper.ParseJSON(t, w, test.ExpectedResponse)
			response := helper.NullifyProperties(responseRaw, []string{"CreatedAt", "UpdatedAt"})

			assert.Equal(t, test.ExpectedResponse, response)
		})
	}
}
