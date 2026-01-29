package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogConfig struct {
	Level      string // debug / info / warn / error
	Format     string // json / console
	OutputPath string // stdout / file（先支持 stdout）
}

var Logger *zap.Logger

func init() {
	logConfig := LogConfig{
		Level:      "info",
		Format:     "json",
		OutputPath: "file",
	}
	InitLogger(logConfig)
}

func InitLogger(cfg LogConfig) {
	level := parseLevel(cfg.Level)

	encoder := getEncoder(cfg.Format)

	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(os.Stdout),
		level,
	)

	Logger = zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)

	zap.ReplaceGlobals(Logger)
}

func parseLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zap.DebugLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

func getEncoder(format string) zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "ts",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stack",
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeTime:    zapcore.ISO8601TimeEncoder,
		EncodeCaller:  zapcore.ShortCallerEncoder,
	}

	if strings.ToLower(format) == "console" {
		return zapcore.NewConsoleEncoder(encoderConfig)
	}

	return zapcore.NewJSONEncoder(encoderConfig)
}
