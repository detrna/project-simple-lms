package pkg

import "time"

type Logger interface {
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	RequestLog(requestID string, method string, path string, statusCode int, start time.Time)
	ErrorLog(method string, path string, statusCode int, err error)
}
