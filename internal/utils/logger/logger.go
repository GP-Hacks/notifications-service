package logger

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	LogLevel         string
	Development      bool
	OutputPaths      []string
	ErrorOutputPaths []string
}

var (
	log  *zap.Logger
	once sync.Once
)

func Initialize(config Config) {
	once.Do(func() {
		if len(config.OutputPaths) == 0 {
			config.OutputPaths = []string{"stdout"}
		}
		if len(config.ErrorOutputPaths) == 0 {
			config.ErrorOutputPaths = []string{"stderr"}
		}

		level := zapcore.InfoLevel
		err := level.UnmarshalText([]byte(config.LogLevel))
		if err != nil {
			level = zapcore.InfoLevel
		}

		encoderConfig := zapcore.EncoderConfig{
			TimeKey:        "timestamp",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}

		zapConfig := zap.Config{
			Level:             zap.NewAtomicLevelAt(level),
			Development:       config.Development,
			Sampling:          nil,
			Encoding:          "json",
			EncoderConfig:     encoderConfig,
			OutputPaths:       config.OutputPaths,
			ErrorOutputPaths:  config.ErrorOutputPaths,
			DisableCaller:     false,
			DisableStacktrace: false,
		}

		if config.Development {
			zapConfig.Encoding = "console"
			zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
			zapConfig.DisableStacktrace = true
			zapConfig.Development = true
		}

		logger, err := zapConfig.Build(
			zap.AddCallerSkip(1),
		)
		if err != nil {
			logger, _ = zap.NewProduction()
		}

		log = logger
	})
}

func Default() {
	Initialize(Config{
		LogLevel:    "info",
		Development: false,
		OutputPaths: []string{"stdout"},
	})
}

func GetLogger() *zap.Logger {
	if log == nil {
		Default()
	}
	return log
}

func With(fields ...zapcore.Field) *zap.Logger {
	return GetLogger().With(fields...)
}

func Named(name string) *zap.Logger {
	return GetLogger().Named(name)
}

func Debug(msg string, fields ...zapcore.Field) {
	GetLogger().Debug(msg, fields...)
}

func Info(msg string, fields ...zapcore.Field) {
	GetLogger().Info(msg, fields...)
}

func Warn(msg string, fields ...zapcore.Field) {
	GetLogger().Warn(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	GetLogger().Error(msg, fields...)
}

func Fatal(msg string, fields ...zapcore.Field) {
	GetLogger().Fatal(msg, fields...)
}

func Sync() error {
	return GetLogger().Sync()
}

func NewRotatingLogger(filename string, maxSize, maxBackups, maxAge int, level zapcore.Level) *zap.Logger {

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)

	core := zapcore.NewCore(
		jsonEncoder,
		zapcore.AddSync(os.Stdout),
		level,
	)

	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}
