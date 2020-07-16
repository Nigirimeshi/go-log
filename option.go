package log

type options struct {
	loggerInstance    loggerInstance
	enableConsole     bool
	consoleJSONFormat bool
	consoleLevel      Level
	enableFile        bool
	fileJSONFormat    bool
	fileLevel         Level
	fileLocation      string
	fileMaxSize       int
	fileMaxBackups    int
	fileMaxAge        int
	fileCompress      bool
}

// Option 保存一个未导出的方法，在一个未导出的 options 结构上记录选项。
type Option interface {
	apply(*options)
}

var _ Option = (*loggerInstanceOption)(nil)

type loggerInstanceOption struct {
	Logger loggerInstance
}

func (l loggerInstanceOption) apply(opts *options) {
	opts.loggerInstance = l.Logger
}

// WithLogger 设置 logger 实例对象, zap 或 logrus
func WithLogger(logger loggerInstance) Option {
	return loggerInstanceOption{
		Logger: logger,
	}
}

var _ Option = (*consoleOption)(nil)

type consoleOption struct {
	Level      Level
	Enable     bool
	JSONFormat bool
}

func (c consoleOption) apply(opts *options) {
	opts.consoleLevel = c.Level
	opts.enableConsole = c.Enable
	opts.consoleJSONFormat = c.JSONFormat
}

// WithConsole 设置控制台相关配置
func WithConsole(enable bool, jsonFormat bool, level Level) Option {
	return consoleOption{
		Enable:     enable,
		JSONFormat: jsonFormat,
		Level:      level,
	}
}

var _ Option = (*fileOption)(nil)

type fileOption struct {
	Level      Level
	Enable     bool
	Location   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

func (f fileOption) apply(opts *options) {
	opts.fileLevel = f.Level
	opts.enableFile = f.Enable
	opts.fileLocation = f.Location
	opts.fileMaxBackups = f.MaxBackups
	opts.fileMaxSize = f.MaxSize
	opts.fileMaxAge = f.MaxAge
	opts.fileCompress = f.Compress
}

// WithFile 设定文件相关配置
func WithFile(
	level Level,
	enable bool,
	location string,
	maxSize int,
	maxBackups int,
	maxAge int,
	compress bool) Option {
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
