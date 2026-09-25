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

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	ts := suite.New()
	sampleData := classfactory.CreateClass(t, ts.DB, "Class-A")
	jwt := authfactory.CreateAdminJWT(t, ts.DB, ts.Infra.TokenService, ts.Infra.Hasher)

	tests := []suite.IntegrationTest[dto.CreateClassRequest]{
		{
			Name: "When create class and success",
			Data: dto.CreateClassRequest{
				SystemID: "new-class-id",
				Name:     sampleData.Name,
			},
			ExpectedStatusCode: http.StatusCreated,
			ExpectedResponse: shared.ResponseSuccess[domain.Class]{
				Data: &domain.Class{
					SystemID: "new-class-id",
					Name:     sampleData.Name,
				},
			},
		},
		{
			Name: "When create class and systemID taken",
			Data: dto.CreateClassRequest{
				SystemID: sampleData.SystemID,
				Name:     sampleData.Name,
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
				http.MethodPost,
				"/api/v1/classes",
				helper.StructToJSON(t, &test.Data),
			)

			req.Header.Set("Authorization", "Bearer "+jwt.Value)

			w := httptest.NewRecorder()

			ts.Router.ServeHTTP(w, req)
			assert.Equal(t, test.ExpectedStatusCode, w.Code)

			responseRaw := helper.ParseJSON(t, w, test.ExpectedResponse)
			response := helper.NullifyProperties(responseRaw, []string{"ID", "CreatedAt", "UpdatedAt"})

			assert.Equal(t, test.ExpectedResponse, response)
		})
	}
}
