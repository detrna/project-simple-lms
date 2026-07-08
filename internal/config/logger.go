package config

import "strconv"

type LoggerConfig struct {
	Level      string
	Format     string
	AddSource  bool
	FilePath   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

func LoadLoggerConfig() *LoggerConfig {
	level := GetEnv("LOG_LEVEL", "info")
	format := GetEnv("LOG_FORMAT", "json")
	addSource, _ := strconv.ParseBool(GetEnv("LOG_ADD_SOURCE", "true"))
	filePath := GetEnv("LOG_FILE_PATH", "logs/app.log")
	maxSize, _ := strconv.Atoi(GetEnv("LOG_MAX_SIZE", "100"))
	maxBackups, _ := strconv.Atoi(GetEnv("LOG_MAX_BACKUPS", "3"))
	maxAge, _ := strconv.Atoi(GetEnv("LOG_MAX_AGE", "28"))
	compress, _ := strconv.ParseBool(GetEnv("LOG_COMPRESS", "true"))

	return &LoggerConfig{
		Level:      level,
		Format:     format,
		AddSource:  addSource,
		FilePath:   filePath,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
	}
}
