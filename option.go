package go_log

type options struct {
	global         bool
	loggerInstance loggerInstance

	enableConsole     bool
	consoleJSONFormat bool
	consoleLevel      Level

	enableFile     bool
	fileJSONFormat bool
	fileLevel      Level
	fileLocation   string
	fileMaxSize    uint // MB
	fileMaxBackups uint
	fileMaxAge     uint // Days
	fileCompress   bool
}

// Option 接口定义了一个未导出的方法，用于覆盖配置。
type Option interface {
	apply(*options)
}

type loggerInstanceOption struct {
	Logger loggerInstance
}

var _ Option = (*loggerInstanceOption)(nil)

func (l loggerInstanceOption) apply(opts *options) {
	opts.loggerInstance = l.Logger
}

// WithLoggerInstance 设置 logger 实例对象。
func WithLoggerInstance(logger loggerInstance) Option {
	return loggerInstanceOption{
		Logger: logger,
	}
}

type consoleOption struct {
	Level      Level
	Enable     bool
	JSONFormat bool
}

var _ Option = (*consoleOption)(nil)

func (c consoleOption) apply(opts *options) {
	opts.consoleLevel = c.Level
	opts.enableConsole = c.Enable
	opts.consoleJSONFormat = c.JSONFormat
}

// WithConsoleOption 设置控制台相关配置。
func WithConsoleOption(
	enable bool,
	jsonFormat bool,
	level Level,
) Option {
	return consoleOption{
		Enable:     enable,
		JSONFormat: jsonFormat,
		Level:      level,
	}
}

type fileOption struct {
	Level      Level
	Enable     bool
	Location   string
	MaxSize    uint
	MaxBackups uint
	MaxAge     uint
	Compress   bool
}

var _ Option = (*fileOption)(nil)

func (f fileOption) apply(opts *options) {
	opts.fileLevel = f.Level
	opts.enableFile = f.Enable
	opts.fileLocation = f.Location
	opts.fileMaxBackups = f.MaxBackups
	opts.fileMaxSize = f.MaxSize
	opts.fileMaxAge = f.MaxAge
	opts.fileCompress = f.Compress
}

// WithFileOption 设定文件相关配置。
func WithFileOption(
	level Level,
	enable bool,
	location string,
	maxSize uint,
	maxBackups uint,
	maxAge uint,
	compress bool,
) Option {
	return fileOption{
		Level:      level,
		Enable:     enable,
		Location:   location,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
	}
}
