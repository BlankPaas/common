package common

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger *zap.SugaredLogger
)

func init() {
	filename := "micro.log"
	syncWriter := zapcore.AddSync(
		&lumberjack.Logger{
			Filename:   filename,
			MaxSize:    512, //mb
			MaxBackups: 0,
			LocalTime:  true,
			Compress:   true,
		},
	)
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoder),
		syncWriter,
		zap.NewAtomicLevelAt(zap.DebugLevel))

	log := zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)

	logger = log.Sugar()
}

func Debug(args ...interface{}) {
	logger.Debug(args)
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	logger.Info(args)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	logger.Warn(args)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	logger.Error(args)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then paniclogger. (See DPanicLevel for detaillogger.)
func DPanic(args ...interface{}) {
	logger.DPanic(args)
}

// Panic uses fmt.Sprint to construct and log a message, then paniclogger.
func Panic(args ...interface{}) {
	logger.Panic(args)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls ologger.Exit.
func Fatal(args ...interface{}) {
	logger.Fatal(args)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then paniclogger. (See DPanicLevel for detaillogger.)
func DPanicf(template string, args ...interface{}) {
	logger.DPanicf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then paniclogger.
func Panicf(template string, args ...interface{}) {
	logger.Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls ologger.Exit.
func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	logger.Fatalf(msg, keysAndValues...)
}
