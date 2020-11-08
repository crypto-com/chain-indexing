package test

import "github.com/crypto-com/chainindex/internal/logger"

type FakeLogger struct{}

func NewFakeLogger() *FakeLogger {
	return &FakeLogger{}
}

func (l *FakeLogger) SetLogLevel(_ logger.LogLevel) {}
func (l *FakeLogger) GetLogLevel() logger.LogLevel {
	return logger.LOG_LEVEL_INFO
}
func (l *FakeLogger) Panic(_ string)                    {}
func (l *FakeLogger) Panicf(_ string, _ ...interface{}) {}
func (l *FakeLogger) Error(_ string)                    {}
func (l *FakeLogger) Errorf(_ string, _ ...interface{}) {}
func (l *FakeLogger) Info(_ string)                     {}
func (l *FakeLogger) Infof(_ string, _ ...interface{})  {}
func (l *FakeLogger) Debug(_ string)                    {}
func (l *FakeLogger) Debugf(_ string, _ ...interface{}) {}

func (l *FakeLogger) WithInterface(_ string, _ interface{}) logger.Logger {
	return l
}

func (l *FakeLogger) WithFields(_ logger.LogFields) logger.Logger {
	return l
}
