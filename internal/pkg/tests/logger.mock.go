package mocks

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Info(msg string, args ...any) {
	m.Called(msg, args)
}

func (m *MockLogger) Warn(msg string, args ...any) {
	m.Called(msg, args)
}

func (m *MockLogger) Error(msg string, args ...any) {
	m.Called(msg, args)
}

func (m *MockLogger) RequestLog(
	requestID string,
	method string,
	path string,
	statusCode int,
	start time.Time,
) {
	m.Called(
		requestID,
		method,
		path,
		statusCode,
		start,
	)
}

func (m *MockLogger) ErrorLog(
	method string,
	path string,
	statusCode int,
	err error,
) {
	m.Called(
		method,
		path,
		statusCode,
		err,
	)
}
