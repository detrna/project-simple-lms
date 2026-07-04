package integration_test

import (
	"encoding/json"
	"fmt"
	"main/internal/app"
	"main/internal/database"
	"main/tests/factory"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetStudentsByClassID(t *testing.T) {
	router := app.NewApp()

	course := factory.CreateCourse(t)

	class := factory.CreateClass(t, course)

	student := factory.CreateUser(t, "Student1")

	factory.EnrollStudent(t, class, student)

	req := httptest.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/api/v1/classes/%s/students", class.ID),
		nil,
	)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var users []database.User

	err := json.Unmarshal(w.Body.Bytes(), &users)
	require.NoError(t, err)

	require.Len(t, users, 1)

	assert.Equal(t, student.ID, users[0].ID)
	assert.Equal(t, student.Name, users[0].Name)
	assert.Equal(t, student.Email, users[0].Email)
}
