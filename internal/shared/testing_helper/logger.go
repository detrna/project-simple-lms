package shared_testing

import (
	pkg_mocks "main/internal/pkg/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
)

func NewMockLogger(t *testing.T) *pkg_mocks.MockLogger {
	mockLogger := pkg_mocks.NewMockLogger(t)
	mockLogger.On("Warn", mock.Anything).Return().Maybe()

	return mockLogger
}
