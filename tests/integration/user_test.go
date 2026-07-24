package integration_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"main/internal/infrastructure/repository/mapper"
	"main/internal/modules/user"
	"main/internal/shared"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUserByID(t *testing.T) {
	exisingUser := Factory.CreateUser(t, "Student1")
	expected := user.UserResponse{
		ID:        exisingUser.ID,
		SystemID:  exisingUser.SystemID,
		Name:      exisingUser.Name,
		Email:     exisingUser.Email,
		Role:      exisingUser.Role,
		CreatedAt: exisingUser.CreatedAt,
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/api/v1/users/%s", exisingUser.ID),
		nil,
	)

	Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response shared.ResponseSuccess[user.UserResponse]
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, expected, *response.Data)
}

func TestCreate(t *testing.T) {
	admin := mapper.ToDomainUser(Factory.CreateAdmin(t))
	token := Factory.CreateJWT(t, admin)

	requestData := user.CreateUserSchema{
		SystemID: "student1",
		Name:     "student1",
		Email:    "student1@mail.com",
		Password: "123",
		Role:     "default",
	}

	time := time.Now()
	expected := user.UserResponse{
		SystemID:  requestData.SystemID,
		Name:      requestData.Name,
		Email:     requestData.Email,
		Role:      requestData.Role,
		CreatedAt: time,
	}

	requestBody, err := json.Marshal(&requestData)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(requestBody))

	req.Header.Set("Authorization", "Bearer "+token.Value)
	req.Header.Set("Content-Type", "application/json")

	Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response shared.ResponseSuccess[user.UserResponse]
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	response.Data.ID = uuid.Nil
	response.Data.CreatedAt = time
	assert.Equal(t, expected, *response.Data)
}

func TestUpdateUser(t *testing.T) {
	admin := mapper.ToDomainUser(Factory.CreateAdmin(t))
	token := Factory.CreateJWT(t, admin)

	existingUser := Factory.CreateUser(t, "Student1")

	newName := "Student2"

	requestData := user.AdminUpdateUserSchema{Name: &newName}

	expected := user.UserResponse{
		ID:        existingUser.ID,
		SystemID:  existingUser.SystemID,
		Name:      newName,
		Email:     existingUser.Email,
		Role:      existingUser.Role,
		CreatedAt: existingUser.CreatedAt,
	}

	requestBody, err := json.Marshal(&requestData)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPatch,
		fmt.Sprintf("/api/v1/users/%s/admin", existingUser.ID),
		bytes.NewBuffer(requestBody),
	)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token.Value)

	Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var user shared.ResponseSuccess[user.UserResponse]
	err = json.Unmarshal(w.Body.Bytes(), &user)
	require.NoError(t, err)

	assert.Equal(t, expected, *user.Data)
}

func TestDeleteUser(t *testing.T) {
	admin := mapper.ToDomainUser(Factory.CreateAdmin(t))
	token := Factory.CreateJWT(t, admin)

	existingUser := Factory.CreateUser(t, "student1")

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/users/%s", existingUser.ID), nil)

	req.Header.Set("Authorization", "Bearer "+token.Value)

	w := httptest.NewRecorder()

	Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
