package helper

import (
	"encoding/json"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func ParseJSON(t *testing.T, w *httptest.ResponseRecorder, expected any) any {
	t.Helper()

	responseType := reflect.TypeOf(expected)
	responsePtr := reflect.New(responseType)

	err := json.Unmarshal(w.Body.Bytes(), responsePtr.Interface())
	require.NoError(t, err)

	return responsePtr.Elem().Interface()
}
