package integration_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"main/internal/infrastructure/database"
	"main/internal/modules/user"
	"main/tests/factory"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUserByID(t *testing.T) {
	indexedUser := factory.CreateUser(t, "Student1")

	req := httptest.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/api/v1/users/%s", indexedUser.ID),
		nil,
	)

	w := httptest.NewRecorder()

	Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var user database.User

	err := json.Unmarshal(w.Body.Bytes(), &user)
	require.NoError(t, err)

	assert.Equal(t, indexedUser.ID, user.ID)
}

func TestCreate(t *testing.T) {
	requestData := user.CreateUserSchema{
		SystemID: "student1",
		Name:     "student1",
		Email:    "student1@mail.com",
		Password: "123",
	}

	requestBody, err := json.Marshal(requestData)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(requestBody))

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var user user.UserResponse

	err = json.Unmarshal(w.Body.Bytes(), &user)
	require.NoError(t, err)

	assert.Equal(t, requestData.SystemID, user.SystemID)
}

func TestUpdateUser(t *testing.T) {
	indexedUser := factory.CreateUser(t, "Student1")

	newName := "Student2"

	requestData := user.UpdateUserSchema{Name: &newName}

	requestBody, err := json.Marshal(requestData)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("/api/v1/users/%s", indexedUser.ID), bytes.NewBuffer(requestBody))

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var user user.UserResponse

	t.Log(w.Body)

	err = json.Unmarshal(w.Body.Bytes(), &user)
	require.NoError(t, err)

	assert.Equal(t, *requestData.Name, user.Name)
}

func TestDeleteUser(t *testing.T) {
	indexedUser := factory.CreateUser(t, "student1")

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/users/%s", indexedUser.ID), nil)

	w := httptest.NewRecorder()

	Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
