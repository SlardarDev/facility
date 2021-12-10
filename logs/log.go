package logs

import (
	"context"
	"fmt"
)

type Logger interface {
	Debugf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Error(format string, v ...interface{})
	Fatalf(format string, v ...interface{})
	Debug(format string, v ...interface{})
	Info(format string, v ...interface{})
	CtxFatal(ctx context.Context, format string, v ...interface{})
	CtxError(ctx context.Context, format string, v ...interface{})
}

type internalLogger struct {
}

func (i *internalLogger) Debugf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func (i *internalLogger) Error(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func (i *internalLogger) Errorf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func (i *internalLogger) Fatalf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func (i *internalLogger) CtxFatal(ctx context.Context, format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func (i *internalLogger) Debug(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func (i *internalLogger) CtxError(ctx context.Context, format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func (i *internalLogger) Info(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

var logger Logger = &internalLogger{}

func SetLogger(log Logger) {
	logger = log
}

func Debugf(format string, v ...interface{}) {
	logger.Debugf(format, v...)
}

func Error(format string, v ...interface{}) {
	logger.Error(format, v...)
}

func Fatalf(format string, v ...interface{}) {
	logger.Fatalf(format, v...)
}

func CtxFatal(ctx context.Context, format string, v ...interface{}) {
	logger.CtxFatal(ctx, format, v...)
}

func Debug(format string, v ...interface{}) {
	logger.Debug(format, v...)
}

func CtxError(ctx context.Context, format string, v ...interface{}) {
	logger.CtxError(ctx, format, v...)
}

func Info(format string, v ...interface{}) {
	logger.Info(format, v...)
}

func Errorf(format string, v ...interface{}) {
	logger.Errorf(format, v...)
}
