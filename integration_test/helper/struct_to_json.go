package helper

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func StructToJSON(t *testing.T, data any) *bytes.Buffer {
	t.Helper()

	json, err := json.Marshal(data)
	require.NoError(t, err)

	return bytes.NewBuffer(json)
}
