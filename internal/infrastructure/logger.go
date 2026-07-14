package infrastructure

import (
	"io"
	"log/slog"
	"main/internal/config"
	"main/internal/pkg"
	"os"
	"strings"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

type SLogger struct {
	*slog.Logger
}

func NewLogger(cfg config.LoggerConfig) pkg.Logger {
	var level slog.Level

	switch strings.ToLower(cfg.Level) {
	case "debug":
		level = slog.LevelDebug
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	writers := []io.Writer{
		os.Stdout,
	}

	writers = append(writers,
		&lumberjack.Logger{
			Filename:   cfg.FilePath,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,
			Compress:   cfg.Compress,
		},
	)

	writer := io.MultiWriter(writers...)

	opts := &slog.HandlerOptions{
		Level:     level,
		AddSource: cfg.AddSource,
	}

	handler := slog.NewJSONHandler(writer, opts)

	return &SLogger{
		Logger: slog.New(handler),
	}
}

func (logger SLogger) RequestLog(requestID string, method string, path string, statusCode int, start time.Time) {
	logger.Info(
		"request completed",
		slog.String("request_id", requestID),
		slog.String("method", method),
		slog.String("path", path),
		slog.Int("status", statusCode),
		slog.Duration("duration", time.Since(start)),
	)
}

func (logger SLogger) ErrorLog(method string, path string, statusCode int, err error) {
	logger.Error(
		"request failed",
		slog.String("method", method),
		slog.String("path", path),
		slog.Int("status", statusCode),
		slog.Any("error", err),
	)
}
