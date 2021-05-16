package go_log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	tests := []struct {
		level Level
		want  string
	}{
		{DEBUG, "debug"},
		{INFO, "info"},
		{WARN, "warn"},
		{ERROR, "error"},
		{FATAL, "fatal"},
		{PANIC, "panic"},
		{6, "undefined"},
	}
	for i := range tests {
		tc := tests[i]
		assert.Equal(t, tc.want, tc.level.String())
	}
}
