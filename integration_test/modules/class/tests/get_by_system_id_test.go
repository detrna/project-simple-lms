package class_test

import (
	"main/integration_test/helper"
	authfactory "main/integration_test/modules/auth/factory"
	classfactory "main/integration_test/modules/class/factory"
	suite "main/integration_test/suite"
	"main/internal/modules/class/domain"
	"main/internal/shared"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBySystemID(t *testing.T) {
	ts := suite.New()
	existingData := classfactory.CreateClass(t, ts.DB, "Class-A")
	jwt := authfactory.CreateAdminJWT(t, ts.DB, ts.Infra.TokenService, ts.Infra.Hasher)

	tests := []suite.IntegrationTest[string]{
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
			ExpectedResponse: shared.ResponseError{
				Error:   domain.ErrClassNotFound.Error(),
				Message: domain.ErrClassNotFound.Message,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			req := httptest.NewRequest(
				http.MethodGet,
				"/api/v1/classes/system/"+test.Data,
				nil,
			)
			req.Header.Set("Authorization", "Bearer "+jwt.Value)

			w := httptest.NewRecorder()

			ts.Router.ServeHTTP(w, req)
			assert.Equal(t, test.ExpectedStatusCode, w.Code)
			req.Header.Set("Authorization", "Bearer "+jwt.Value)

			response := helper.ParseJSON(t, w, test.ExpectedResponse)

			assert.Equal(t, test.ExpectedResponse, response)
		})
	}
}
