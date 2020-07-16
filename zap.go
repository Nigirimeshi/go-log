package log

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var _ Logger = (*zapLogger)(nil)

type zapLogger struct {
	sugaredLogger *zap.SugaredLogger
	callerSkipFix bool
}

func newZapLogger(opts options) (Logger, error) {
	cores := []zapcore.Core{}

	if opts.enableConsole {
		level := getZapLevel(opts.consoleLevel)
		writer := zapcore.Lock(os.Stdout)
		core := zapcore.NewCore(getEncoder(opts.consoleJSONFormat), writer, level)
		cores = append(cores, core)
	}

	if opts.enableFile {
		level := getZapLevel(opts.fileLevel)
		writer := zapcore.AddSync(&lumberjack.Logger{
			Filename: opts.fileLocation,
			MaxSize:  opts.fileMaxSize,
			Compress: opts.fileCompress,
			MaxAge:   opts.fileMaxAge,
		})
		core := zapcore.NewCore(getEncoder(opts.fileJSONFormat), writer, level)
		cores = append(cores, core)
	}

	combinedCore := zapcore.NewTee(cores...)

	logger := zap.New(combinedCore,
		zap.AddCallerSkip(2),
		zap.AddCaller(),
	)

	zap.RedirectStdLogAt(logger, zap.InfoLevel)

	return &zapLogger{
		sugaredLogger: logger.Sugar(),
	}, nil
}

func (l *zapLogger) Debugf(format string, args ...interface{}) {
	l.sugaredLogger.Debugf(format, args...)
}

func (l *zapLogger) Infof(format string, args ...interface{}) {
	l.sugaredLogger.Infof(format, args...)
}

func (l *zapLogger) Warnf(format string, args ...interface{}) {
	l.sugaredLogger.Warnf(format, args...)
}

func (l *zapLogger) Errorf(format string, args ...interface{}) {
	l.sugaredLogger.Errorf(format, args...)
}

func (l *zapLogger) Fatalf(format string, args ...interface{}) {
	l.sugaredLogger.Fatalf(format, args...)
}

func (l *zapLogger) Panicf(format string, args ...interface{}) {
	l.sugaredLogger.Panicf(format, args...)
}

func (l *zapLogger) WithFields(fields Fields) Logger {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}

	var newLogger *zap.SugaredLogger

	if !l.callerSkipFix {
		newLogger = l.sugaredLogger.Desugar().
			WithOptions(zap.AddCallerSkip(-1)).
			Sugar().With(f...)
	} else {
		newLogger = l.sugaredLogger.With(f...)
	}

	return &zapLogger{newLogger, true}
}

func getEncoder(isJSON bool) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = getTimeEncoder
	encoderConfig.TimeKey = "time"
	if isJSON {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(timestampFormat))
	// enc.AppendString(t.Format(time.RFC822Z))
}

func getZapLevel(level Level) zapcore.Level {
	switch level {
	case DEBUG:
		return zapcore.DebugLevel
	case INFO:
		return zapcore.InfoLevel
	case WARN:
		return zapcore.WarnLevel
	case ERROR:
		return zapcore.ErrorLevel
	case FATAL:
		return zapcore.FatalLevel
	case PANIC:
		return zapcore.PanicLevel
	default:
		return zapcore.InfoLevel
	}
}
