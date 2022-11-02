package infrastructure

import (
	"fmt"
	"io"
	"time"

	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/rs/zerolog"
)

// ZerologLogger implements applogger.Logger interface and provide logging service
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

func (logger *ZerologLogger) SetLogLevel(level applogger.LogLevel) {
	updatedLogger := logger.instance.Level(logLevelToZerologLevel(level))

	logger.instance = &updatedLogger
}

func logLevelToZerologLevel(logLevel applogger.LogLevel) zerolog.Level {
	switch logLevel {
	case applogger.LOG_DISABLED:
		return zerolog.Disabled
	case applogger.LOG_LEVEL_PANIC:
		return zerolog.PanicLevel
	case applogger.LOG_LEVEL_ERROR:
		return zerolog.ErrorLevel
	case applogger.LOG_LEVEL_WARN:
		return zerolog.WarnLevel
	case applogger.LOG_LEVEL_INFO:
		return zerolog.InfoLevel
	case applogger.LOG_LEVEL_DEBUG:
		return zerolog.DebugLevel
	default:
		panic(fmt.Sprintf("Unsupported log level %v", logLevel))
	}
}

func (logger *ZerologLogger) GetLogLevel() applogger.LogLevel {
	return zerologLevelToLogLevel(logger.instance.GetLevel())
}

func zerologLevelToLogLevel(logLevel zerolog.Level) applogger.LogLevel {
	switch logLevel {
	case zerolog.Disabled:
		return applogger.LOG_DISABLED
	case zerolog.PanicLevel:
	case zerolog.FatalLevel:
		return applogger.LOG_LEVEL_PANIC
	case zerolog.ErrorLevel:
		return applogger.LOG_LEVEL_ERROR
	case zerolog.WarnLevel:
		return applogger.LOG_LEVEL_WARN
	case zerolog.InfoLevel:
		return applogger.LOG_LEVEL_INFO
	case zerolog.DebugLevel:
	case zerolog.TraceLevel:
	case zerolog.NoLevel:
		return applogger.LOG_LEVEL_DEBUG
	default:
		panic(fmt.Sprintf("Unsupported log level %v", logLevel))
	}

	return applogger.LOG_LEVEL_DEBUG
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

func (logger *ZerologLogger) Warn(message string) {
	logger.instance.Warn().Timestamp().Msg(message)
}

func (logger *ZerologLogger) Warnf(format string, values ...interface{}) {
	logger.instance.Warn().Timestamp().Msgf(format, values...)
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

func (logger *ZerologLogger) WithInterface(key string, value interface{}) applogger.Logger {
	instance := logger.instance.With().Interface(key, value).Logger()
	return &ZerologLogger{
		instance: &instance,
	}
}

func (logger *ZerologLogger) WithFields(fields applogger.LogFields) applogger.Logger {
	instance := logger.instance.With().Fields(fields).Logger()
	return &ZerologLogger{
		instance: &instance,
	}
}
