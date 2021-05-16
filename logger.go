package go_log

type loggerInstance uint8

const (
	// Zap 表示 zap 实例。
	Zap loggerInstance = iota
	// Logrus 表示 logrus 实例。
	Logrus
	// DefaultLogger 表示默认的日志实例。
	DefaultLogger = Zap
)

// Fields Type to pass when we want to call WithFields for structured logging.
type Fields map[string]interface{}

// Logger 定义了日志接口。
type Logger interface {
	Debugf(format string, args ...interface{})

	Infof(format string, args ...interface{})

	Warnf(format string, args ...interface{})

	Errorf(format string, args ...interface{})

	Fatalf(format string, args ...interface{})

	Panicf(format string, args ...interface{})

	WithFields(fields Fields) Logger
}

// New 返回一个 Logger 接口实例。
func New(opts ...Option) (Logger, error) {
	options := options{
		global:         true,
		loggerInstance: DefaultLogger,

		enableConsole:     true,
		consoleJSONFormat: false,
		consoleLevel:      DEFAULT_LEVEL,

		enableFile:     false,
		fileJSONFormat: false,
		fileLevel:      DEFAULT_LEVEL,
		fileLocation:   "",
		fileMaxSize:    100,
		fileMaxBackups: 5,
		fileMaxAge:     7,
		fileCompress:   false,
	}
	for _, o := range opts {
		o.apply(&options)
	}

	var logger Logger
	switch options.loggerInstance {
	case Zap:
		zapLogger, err := newZapLogger(options)
		if err != nil {
			return nil, err
		}
		logger = zapLogger
		return logger, err

	case Logrus:
		logrusLogger, err := newLogrusLogger(options)
		if err != nil {
			return nil, err
		}
		logger = logrusLogger
		return logger, nil

	default:
		return nil, errInvalidLoggerInstance
	}
}

// _logger 是默认的全局日志。
var _logger Logger

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
