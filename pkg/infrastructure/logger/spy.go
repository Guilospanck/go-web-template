package logger

import "go.uber.org/zap"

type loggerSpy struct{}

func (spy loggerSpy) Debug(msg string, fields ...zap.Field) {

}
func (spy loggerSpy) Info(msg string, fields ...zap.Field) {

}
func (spy loggerSpy) Warn(msg string, fields ...zap.Field) {

}
func (spy loggerSpy) Error(msg string, fields ...zap.Field) {

}
func (spy loggerSpy) Fatal(msg string, fields ...zap.Field) {

}

func NewLoggerSpy() loggerSpy {
	return loggerSpy{}
}
