package log

import (
	"testing"
)

func TestNew(t *testing.T) {
	if err := New(); err != nil {
		panic(err)
	}

	Debugf("debug...")
	Infof("info...")
	Warnf("warn...")
	Errorf("error...")

	contextFieldsLogger := WithFields(Fields{"key1": "value1"})
	contextFieldsLogger.Warnf("content warn...")
	// contextFieldsLogger.Panicf("content fatal...")
}
