package fake

import "github.com/crypto-com/chainindex/usecase"

type FakeLogger struct{}

func NewFakeLogger() *FakeLogger {
	return &FakeLogger{}
}

func (logger *FakeLogger) SetLogLevel(_ usecase.LogLevel) {}
func (logger *FakeLogger) GetLogLevel() usecase.LogLevel {
	return usecase.LOG_LEVEL_INFO
}
func (logger *FakeLogger) Panic(_ string)                    {}
func (logger *FakeLogger) Panicf(_ string, _ ...interface{}) {}
func (logger *FakeLogger) Error(_ string)                    {}
func (logger *FakeLogger) Errorf(_ string, _ ...interface{}) {}
func (logger *FakeLogger) Info(_ string)                     {}
func (logger *FakeLogger) Infof(_ string, _ ...interface{})  {}
func (logger *FakeLogger) Debug(_ string)                    {}
func (logger *FakeLogger) Debugf(_ string, _ ...interface{}) {}

func (logger *FakeLogger) WithInterface(_ string, _ interface{}) usecase.Logger {
	return logger
}

func (logger *FakeLogger) WithFields(_ usecase.LogFields) usecase.Logger {
	return logger
}
