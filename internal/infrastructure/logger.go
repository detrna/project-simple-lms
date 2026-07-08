package infrastructure

import (
	"io"
	"log/slog"
	"main/internal/config"
	"os"
	"strings"

	"gopkg.in/natefinch/lumberjack.v2"
)

type SLogger struct {
	*slog.Logger
	config config.LoggerConfig
}

func NewLogger(cfg config.LoggerConfig) *slog.Logger {
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

	return slog.New(handler)
}
