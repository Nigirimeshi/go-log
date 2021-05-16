package go_log

const (
	// DEBUG ...
	DEBUG Level = iota
	// INFO ...
	INFO
	// WARN ...
	WARN
	// ERROR ...
	ERROR
	// FATAL 用于记录致命错误信息，记录后进程会关闭。
	FATAL
	// PANIC ...
	PANIC
	// DEFAULT_LEVEL ...
	DEFAULT_LEVEL = INFO
)

// Level 表示日志级别。
type Level uint8

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
		return "undefined"
	}
}
