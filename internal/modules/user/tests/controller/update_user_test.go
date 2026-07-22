package user_controller_test

import (
	"bytes"
	"context"
	"encoding/json"
	"main/internal/modules/user"
	user_mocks "main/internal/modules/user/mocks"
	user_factory "main/internal/modules/user/tests"
	"main/internal/shared"
	shared_testing "main/internal/shared/testing_helper"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUpdateUser_Success(t *testing.T) {
	mockUseCase := user_mocks.NewMockIUseCase(t)
	mockLogger := shared_testing.NewMockLogger(t)

	id := uuid.New()
	existingAccount := user_factory.NewUser(id)
	jwtPayload := user_factory.NewJWTPayload(existingAccount)

	newPassword := "password321"
	requestData := user.UpdateUserSchema{
		Password: &newPassword,
	}

	usecaseResult := user.UserResponse{
		ID:       existingAccount.ID,
		SystemID: existingAccount.SystemID,
		Name:     existingAccount.Name,
		Email:    existingAccount.Email,
		Role:     existingAccount.Role,
	}

	expected := shared.ResponseSuccess[user.UserResponse]{
		Data: &usecaseResult,
	}

	ctx := context.Background()
	mockUseCase.On("UpdateUser", ctx, mock.AnythingOfType("*user.UpdateUserDTO")).Return(&usecaseResult, nil)

	c := user.NewController(mockUseCase, mockLogger)
	router := gin.New()

	router.PATCH("/:id", func(ctx *gin.Context) {
		ctx.Set("user", jwtPayload)
		c.UpdateUser(ctx)
	})

	w := httptest.NewRecorder()

	requestBody, err := json.Marshal(&requestData)
	require.NoError(t, err)

	req := httptest.NewRequest(
		http.MethodPatch,
		("/" + id.String()),
		bytes.NewReader(requestBody),
	)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response shared.ResponseSuccess[user.UserResponse]
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, expected.Data, response.Data)

	mockUseCase.AssertExpectations(t)
}
