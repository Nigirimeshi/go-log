package go_log

import (
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var _ Logger = (*logrusLogEntry)(nil)
var _ Logger = (*logrusLogger)(nil)

type logrusLogEntry struct {
	entry *logrus.Entry
}

type logrusLogger struct {
	logger *logrus.Logger
}

func newLogrusLogger(opts options) (Logger, error) {
	logLevel := opts.consoleLevel

	level, err := logrus.ParseLevel(logLevel.String())
	if err != nil {
		return nil, err
	}

	stdOutHandler := os.Stdout
	fileHandler := &lumberjack.Logger{
		Filename: opts.fileLocation,
		MaxSize:  int(opts.fileMaxSize),
		Compress: opts.fileCompress,
		MaxAge:   int(opts.fileMaxAge),
	}
	lLogger := &logrus.Logger{
		Out:          stdOutHandler,
		Formatter:    getFormatter(opts.consoleJSONFormat),
		Hooks:        make(logrus.LevelHooks),
		Level:        level,
		ReportCaller: true,
	}

	if opts.enableConsole && opts.enableFile {
		// 同时启用控制台和文件记录时，以文件记录的格式为准
		lLogger.SetOutput(io.MultiWriter(stdOutHandler, fileHandler))
		lLogger.SetFormatter(getFormatter(opts.fileJSONFormat))
	} else {
		if opts.enableFile {
			lLogger.SetOutput(fileHandler)
			lLogger.SetFormatter(getFormatter(opts.fileJSONFormat))
		}
	}

	return &logrusLogger{
		logger: lLogger,
	}, nil
}

func getFormatter(isJSON bool) logrus.Formatter {
	if isJSON {
		return &logrus.JSONFormatter{
			TimestampFormat: time.RFC3339Nano,
		}
	}
	return &logrus.TextFormatter{
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		TimestampFormat:        time.RFC3339Nano,
	}
}

func (l *logrusLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *logrusLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *logrusLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *logrusLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *logrusLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *logrusLogger) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(format, args...)
}

func (l *logrusLogger) WithFields(fields Fields) Logger {
	return &logrusLogEntry{
		entry: l.logger.WithFields(convertToLogrusFields(fields)),
	}
}

func (l *logrusLogEntry) Debugf(format string, args ...interface{}) {
	l.entry.Debugf(format, args...)
}

func (l *logrusLogEntry) Infof(format string, args ...interface{}) {
	l.entry.Infof(format, args...)
}

func (l *logrusLogEntry) Warnf(format string, args ...interface{}) {
	l.entry.Warnf(format, args...)
}

func (l *logrusLogEntry) Errorf(format string, args ...interface{}) {
	l.entry.Errorf(format, args...)
}

func (l *logrusLogEntry) Fatalf(format string, args ...interface{}) {
	l.entry.Fatalf(format, args...)
}

func (l *logrusLogEntry) Panicf(format string, args ...interface{}) {
	l.entry.Panicf(format, args...)
}

func (l *logrusLogEntry) WithFields(fields Fields) Logger {
	return &logrusLogEntry{
		entry: l.entry.WithFields(convertToLogrusFields(fields)),
	}
}

func convertToLogrusFields(fields Fields) logrus.Fields {
	logrusFields := logrus.Fields{}
	for index, val := range fields {
		logrusFields[index] = val
	}
	return logrusFields
}
