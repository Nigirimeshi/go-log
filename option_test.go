package log

import "testing"

func TestWithLogger(t *testing.T) {
	if err := New(WithLogger(InstanceLogrusLogger)); err != nil {
		panic(err)
	}
	Debugf("debug...")
	Infof("info...")
	Warnf("warn...")
	Errorf("error...")
}
