package log

import "errors"

var _logger Logger

// Level represents a log level.
type Level int32

const (
	// DEBUG information for programmer low level analysis.
	DEBUG Level = iota + 1
	// INFO information about steady state operations.
	INFO
	// WARN is for logging messages about possible issues.
	WARN
	// ERROR is for logging errors.
	ERROR
	// FATAL is for logging fatal messages. The system shutdown after logging the message.
	FATAL
	// PANIC ...
	PANIC
)

type loggerInstance int

const (
	// InstanceZapLogger is zap logger instance.
	InstanceZapLogger loggerInstance = iota + 1
	// InstanceLogrusLogger is logrus logger instance.
	InstanceLogrusLogger
	// DefaultLogger is zap logger instance.
	DefaultLogger = InstanceZapLogger
)

const (
	timestampFormat = "2006-01-02 15:04:05 -0700"
)

var (
	errInvalidLoggerInstance = errors.New("invalid logger instance")
)

// Fields Type to pass when we want to call WithFields for structured logging.
type Fields map[string]interface{}

// Logger ...
type Logger interface {
	Debugf(format string, args ...interface{})

	Infof(format string, args ...interface{})

	Warnf(format string, args ...interface{})

	Errorf(format string, args ...interface{})

	Fatalf(format string, args ...interface{})

	Panicf(format string, args ...interface{})

	WithFields(fields Fields) Logger
}

func (l Level) String() string {
	switch l {
	case DEBUG:
		return "debug"
	case INFO:
		return "info"
	case WARN:
		return "warn"
	case ERROR:
		return "error"
	case FATAL:
		return "fatal"
	case PANIC:
		return "panic"
	default:
		return "unknown"
	}
}

// New returns an instance of logger
func New(opts ...Option) error {
	options := options{
		loggerInstance:    DefaultLogger,
		enableConsole:     true,
		consoleJSONFormat: false,
		consoleLevel:      INFO,
		enableFile:        false,
		fileJSONFormat:    false,
		fileLevel:         INFO,
		fileLocation:      "",
		fileMaxSize:       100,
		fileMaxBackups:    3,
		fileMaxAge:        7,
		fileCompress:      false,
	}
	for _, o := range opts {
		o.apply(&options)
	}
	switch options.loggerInstance {
	case InstanceZapLogger:
		zapLogger, err := newZapLogger(options)
		if err != nil {
			return err
		}
		_logger = zapLogger
		return nil

	case InstanceLogrusLogger:
		logrusLogger, err := newLogrusLogger(options)
		if err != nil {
			return err
		}
		_logger = logrusLogger
		return nil

	default:
		return errInvalidLoggerInstance
	}
}

// Debugf ...
func Debugf(format string, args ...interface{}) {
	_logger.Debugf(format, args...)
}

// Infof ...
func Infof(format string, args ...interface{}) {
	_logger.Infof(format, args...)
}

// Warnf ...
func Warnf(format string, args ...interface{}) {
	_logger.Warnf(format, args...)
}

// Errorf ...
func Errorf(format string, args ...interface{}) {
	_logger.Errorf(format, args...)
}

// Fatalf ...
func Fatalf(format string, args ...interface{}) {
	_logger.Fatalf(format, args...)
}

// Panicf ...
func Panicf(format string, args ...interface{}) {
	_logger.Panicf(format, args...)
}

// WithFields ...
func WithFields(fileds Fields) Logger {
	return _logger.WithFields(fileds)
}
