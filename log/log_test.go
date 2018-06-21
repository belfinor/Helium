package log

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2017-06-21

import (
	"testing"
)

func TestLoggerSetLevel(t *testing.T) {

	defLog = &Log{}

	if GetLevel() != "none" {
		t.Fatal("invalid default log level")
	}

	SetLevel("info")

	if GetLevel() != "info" {
		t.Fatal("expected log level 'info'")
	}

	defLog = nil

	Error("error1", "error2")
	Warn("warn")
	Info("info")
	Debug("debug")
	Trace("trace", 1, "trace", 2)
}
