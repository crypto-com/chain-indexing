package infrastructure

import (
	"fmt"
	"io"
	"time"

	"github.com/rs/zerolog"

	"github.com/crypto-com/chainindex/usecase"
)

// ZerologLogger implements usecase.Logger interface and provide logging service
// using zerolog.
type ZerologLogger struct {
	instance *zerolog.Logger
}

func NewZerologLogger(output io.Writer) *ZerologLogger {
	logger := zerolog.New(zerolog.ConsoleWriter{
		Out:        output,
		NoColor:    true,
		TimeFormat: time.RFC3339,
	})

	return &ZerologLogger{
		instance: &logger,
	}
}

func NewZerologLoggerWithColor(output io.Writer) *ZerologLogger {
	logger := zerolog.New(zerolog.ConsoleWriter{
		Out:        output,
		NoColor:    false,
		TimeFormat: time.RFC3339,
	})

	return &ZerologLogger{
		instance: &logger,
	}
}

func (logger *ZerologLogger) SetLogLevel(level usecase.LogLevel) {
	updatedLogger := logger.instance.Level(logLevelToZerologLevel(level))

	logger.instance = &updatedLogger
}

func logLevelToZerologLevel(logLevel usecase.LogLevel) zerolog.Level {
	switch logLevel {
	case usecase.LOG_DISABLED:
		return zerolog.Disabled
	case usecase.LOG_LEVEL_PANIC:
		return zerolog.PanicLevel
	case usecase.LOG_LEVEL_ERROR:
		return zerolog.ErrorLevel
	case usecase.LOG_LEVEL_INFO:
		return zerolog.InfoLevel
	case usecase.LOG_LEVEL_DEBUG:
		return zerolog.DebugLevel
	default:
		panic(fmt.Sprintf("Unsupported log level %v", logLevel))
	}
}

func (logger *ZerologLogger) GetLogLevel() usecase.LogLevel {
	return zerologLevelToLogLevel(logger.instance.GetLevel())
}

func zerologLevelToLogLevel(logLevel zerolog.Level) usecase.LogLevel {
	switch logLevel {
	case zerolog.Disabled:
		return usecase.LOG_DISABLED
	case zerolog.PanicLevel:
		return usecase.LOG_LEVEL_PANIC
	case zerolog.ErrorLevel:
		return usecase.LOG_LEVEL_ERROR
	case zerolog.InfoLevel:
		return usecase.LOG_LEVEL_INFO
	case zerolog.DebugLevel:
		return usecase.LOG_LEVEL_DEBUG
	default:
		panic(fmt.Sprintf("Unsupported log level %v", logLevel))
	}
}

func (logger *ZerologLogger) Panic(message string) {
	logger.instance.Panic().Timestamp().Msg(message)
}

func (logger *ZerologLogger) Panicf(format string, values ...interface{}) {
	logger.instance.Panic().Timestamp().Msgf(format, values...)
}

func (logger *ZerologLogger) Error(message string) {
	logger.instance.Error().Timestamp().Msg(message)
}

func (logger *ZerologLogger) Errorf(format string, values ...interface{}) {
	logger.instance.Error().Timestamp().Msgf(format, values...)
}

func (logger *ZerologLogger) Info(message string) {
	logger.instance.Info().Timestamp().Msg(message)
}

func (logger *ZerologLogger) Infof(format string, values ...interface{}) {
	logger.instance.Info().Timestamp().Msgf(format, values...)
}

func (logger *ZerologLogger) Debug(message string) {
	logger.instance.Debug().Timestamp().Msg(message)
}

func (logger *ZerologLogger) Debugf(format string, values ...interface{}) {
	logger.instance.Debug().Timestamp().Msgf(format, values...)
}

func (logger *ZerologLogger) WithInterface(key string, value interface{}) usecase.Logger {
	instance := logger.instance.With().Interface(key, value).Logger()
	return &ZerologLogger{
		instance: &instance,
	}
}

func (logger *ZerologLogger) WithFields(fields usecase.LogFields) usecase.Logger {
	instance := logger.instance.With().Fields(fields).Logger()
	return &ZerologLogger{
		instance: &instance,
	}
}
