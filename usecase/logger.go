package usecase

type Logger interface {
	SetLogLevel(level LogLevel)
	GetLogLevel() LogLevel

	Panic(message string)
	Panicf(format string, values ...interface{})
	Error(message string)
	Errorf(format string, values ...interface{})
	Info(message string)
	Infof(format string, values ...interface{})
	Debug(message string)
	Debugf(format string, values ...interface{})

	WithInterface(key string, value interface{}) Logger
	WithFields(fields LogFields) Logger
}

type LogLevel int8

const (
	LOG_LEVEL_DEBUG LogLevel = iota
	LOG_LEVEL_INFO
	LOG_LEVEL_ERROR
	LOG_LEVEL_PANIC
	LOG_DISABLED
)

type LogFields map[string]interface{}
