package go_log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithLoggerInstance(t *testing.T) {
	tests := []struct {
		loggerInstance loggerInstance
	}{
		{Zap},
		{Logrus},
	}
	for i := range tests {
		tc := tests[i]
		opts := options{}
		o := WithLoggerInstance(tc.loggerInstance)
		o.apply(&opts)
		assert.Equal(t, tc.loggerInstance, opts.loggerInstance)
	}
}

func TestWithConsoleOption(t *testing.T) {
	tests := []struct {
		enable     bool
		jsonFormat bool
		level      Level
	}{
		{false, false, DEBUG},
		{false, true, INFO},
		{true, true, WARN},
	}
	for i := range tests {
		tc := tests[i]
		opts := options{}
		o := WithConsoleOption(tc.enable, tc.jsonFormat, tc.level)
		o.apply(&opts)
		assert.Equal(t, tc.enable, opts.enableConsole)
		assert.Equal(t, tc.jsonFormat, opts.consoleJSONFormat)
		assert.Equal(t, tc.level, opts.consoleLevel)
	}
}

func TestWithFileOption(t *testing.T) {
	tests := []struct {
		level      Level
		enable     bool
		location   string
		maxSize    uint
		maxBackups uint
		maxAge     uint
		compress   bool
	}{
		{DEBUG, false, "", 0, 0, 0, false},
		{INFO, false, "", 0, 0, 0, true},
		{WARN, false, "", 0, 0, 1, true},
		{ERROR, false, "", 0, 1, 2, true},
		{FATAL, false, "", 1, 2, 3, true},
		{FATAL, false, "/a/b/c", 1, 2, 3, true},
		{PANIC, true, "/a/b/c", 1, 2, 3, true},
	}
	for i := range tests {
		tc := tests[i]
		opts := options{}
		o := WithFileOption(tc.level, tc.enable, tc.location, tc.maxSize, tc.maxBackups, tc.maxAge, tc.compress)
		o.apply(&opts)
		assert.Equal(t, tc.level, opts.fileLevel)
		assert.Equal(t, tc.enable, opts.enableFile)
		assert.Equal(t, tc.location, opts.fileLocation)
		assert.Equal(t, tc.maxSize, opts.fileMaxSize)
		assert.Equal(t, tc.maxBackups, opts.fileMaxBackups)
		assert.Equal(t, tc.maxAge, opts.fileMaxAge)
		assert.Equal(t, tc.compress, opts.fileCompress)
	}
}
