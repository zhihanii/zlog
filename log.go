package zlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

var (
	logger = newZapLogger(NewOptions())
	mu     sync.Mutex
)

type Logger interface {
	Debug(msg string, fields ...Field)
	Debugf(format string, v ...any)
	Debugw(msg string, keysAndValues ...any)
	Info(msg string, fields ...Field)
	Infof(format string, v ...any)
	Infow(msg string, keysAndValues ...any)
	Warn(msg string, fields ...Field)
	Warnf(format string, v ...any)
	Warnw(msg string, keysAndValues ...any)
	Error(msg string, fields ...Field)
	Errorf(format string, v ...any)
	Errorw(msg string, keysAndValues ...any)
	Panic(msg string, fields ...Field)
	Panicf(format string, v ...any)
	Panicw(msg string, keysAndValues ...any)
	Fatal(msg string, fields ...Field)
	Fatalf(format string, v ...any)
	Fatalw(msg string, keysAndValues ...any)
}

var _ Logger = &zapLogger{}

// Init initializes logger with specified options.
func Init(o *Options) {
	mu.Lock()
	defer mu.Unlock()
	logger = newZapLogger(o)
}

func New(o *Options) Logger {
	return newZapLogger(o)
}

func Debug(msg string, fields ...Field) {
	logger.zl.Debug(msg, fields...)
}

func Debugf(format string, v ...any) {
	logger.sl.Debugf(format, v...)
}

func Debugw(msg string, keysAndValues ...any) {
	logger.sl.Debugw(msg, keysAndValues...)
}

func Debugln(args ...any) {
	logger.sl.Debugln(args...)
}

func Info(msg string, fields ...Field) {
	logger.zl.Info(msg, fields...)
}

func Infof(format string, v ...any) {
	logger.sl.Infof(format, v...)
}

func Infow(msg string, keysAndValues ...any) {
	logger.sl.Infow(msg, keysAndValues...)
}

func Infoln(args ...any) {
	logger.sl.Infoln(args...)
}

func Warn(msg string, fields ...Field) {
	logger.zl.Warn(msg, fields...)
}

func Warnf(format string, v ...any) {
	logger.sl.Warnf(format, v...)
}

func Warnw(msg string, keysAndValues ...any) {
	logger.sl.Warnw(msg, keysAndValues...)
}

func Warnln(args ...any) {
	logger.sl.Warnln(args...)
}

func Error(msg string, fields ...Field) {
	logger.zl.Error(msg, fields...)
}

func Errorf(format string, v ...any) {
	logger.sl.Errorf(format, v...)
}

func Errorw(msg string, keysAndValues ...any) {
	logger.sl.Errorw(msg, keysAndValues...)
}

func Errorln(args ...any) {
	logger.sl.Errorln(args...)
}

func Panic(msg string, fields ...Field) {
	logger.zl.Panic(msg, fields...)
}

func Panicf(format string, v ...any) {
	logger.sl.Panicf(format, v...)
}

func Panicw(msg string, keysAndValues ...any) {
	logger.sl.Panicw(msg, keysAndValues...)
}

func Panicln(args ...any) {
	logger.sl.Panicln(args...)
}

func Fatal(msg string, fields ...Field) {
	logger.zl.Fatal(msg, fields...)
}

func Fatalf(format string, v ...any) {
	logger.sl.Fatalf(format, v...)
}

func Fatalw(msg string, keysAndValues ...any) {
	logger.sl.Fatalw(msg, keysAndValues...)
}

func Fatalln(args ...any) {
	logger.sl.Fatalln(args...)
}

type zapLogger struct {
	zl *zap.Logger
	sl *zap.SugaredLogger
}

func newZapLogger(o *Options) *zapLogger {
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(o.Level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}
	encodeLevel := zapcore.CapitalLevelEncoder
	// when output to local path, with color is forbidden
	if o.Format == consoleFormat && o.EnableColor {
		encodeLevel = zapcore.CapitalColorLevelEncoder
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "timestamp",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    encodeLevel,
		EncodeTime:     timeEncoder,
		EncodeDuration: milliSecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	loggerConfig := &zap.Config{
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Development:       o.Development,
		DisableCaller:     o.DisableCaller,
		DisableStacktrace: o.DisableStacktrace,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         o.Format,
		EncoderConfig:    encoderConfig,
		OutputPaths:      o.OutputPaths,
		ErrorOutputPaths: o.ErrorOutputPaths,
	}

	var err error
	l, err := loggerConfig.Build(zap.AddStacktrace(zapcore.PanicLevel), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	z := &zapLogger{
		zl: l.Named(o.Name),
		sl: l.Sugar(),
	}

	return z
}

func (l *zapLogger) Debug(msg string, fields ...Field) {
	l.zl.Debug(msg, fields...)
}

func (l *zapLogger) Debugf(format string, v ...any) {
	l.sl.Debugf(format, v...)
}

func (l *zapLogger) Debugw(msg string, keysAndValues ...any) {
	l.sl.Debugw(msg, keysAndValues...)
}

func (l *zapLogger) Info(msg string, fields ...Field) {
	l.zl.Info(msg, fields...)
}

func (l *zapLogger) Infof(format string, v ...any) {
	l.sl.Infof(format, v...)
}

func (l *zapLogger) Infow(msg string, keysAndValues ...any) {
	l.sl.Infow(msg, keysAndValues...)
}

func (l *zapLogger) Warn(msg string, fields ...Field) {
	l.zl.Warn(msg, fields...)
}

func (l *zapLogger) Warnf(format string, v ...any) {
	l.sl.Warnf(format, v...)
}

func (l *zapLogger) Warnw(msg string, keysAndValues ...any) {
	l.sl.Warnw(msg, keysAndValues...)
}

func (l *zapLogger) Error(msg string, fields ...Field) {
	l.zl.Error(msg, fields...)
}

func (l *zapLogger) Errorf(format string, v ...any) {
	l.sl.Errorf(format, v...)
}

func (l *zapLogger) Errorw(msg string, keysAndValues ...any) {
	l.sl.Errorw(msg, keysAndValues...)
}

func (l *zapLogger) Panic(msg string, fields ...Field) {
	l.zl.Panic(msg, fields...)
}

func (l *zapLogger) Panicf(format string, v ...any) {
	l.sl.Panicf(format, v...)
}

func (l *zapLogger) Panicw(msg string, keysAndValues ...any) {
	l.sl.Panicw(msg, keysAndValues...)
}

func (l *zapLogger) Fatal(msg string, fields ...Field) {
	l.zl.Fatal(msg, fields...)
}

func (l *zapLogger) Fatalf(format string, v ...any) {
	l.sl.Fatalf(format, v...)
}

func (l *zapLogger) Fatalw(msg string, keysAndValues ...any) {
	l.sl.Fatalw(msg, keysAndValues...)
}
