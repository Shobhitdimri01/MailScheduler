package logs

import (
	"os"
	"strings"
	"testing"
)

func TestNewLogger(t *testing.T) {
	testlogpath := "logs/test.log"
	os.Remove(testlogpath)
	defer func() {
		if err := os.RemoveAll("logs"); err != nil {
			t.Error("Error in removing log path err:" + err.Error())
		}
	}()
	l, err := NewLogger(testlogpath)
	if err != nil {
		t.Error("Failed initializing new logger")
	}
	infomsg := "test log entry"
	l.Info(infomsg)
	l.Warn("This is a warning")
	l.Error("This is an error")
	l.Close()
	_, err = os.Stat(testlogpath)
	if err != nil {
		t.Error("log file not found! - " + err.Error())
	}

	content, readerr := os.ReadFile(testlogpath)
	if readerr != nil {
		t.Error("log content not found - " + err.Error())
	}
	if !strings.Contains(string(content), "info") {
		t.Error("info msg not found ")
	}
	if !strings.Contains(string(content), "error") {
		t.Error("error msg not found ")
	}
	if !strings.Contains(string(content), "warn") {
		t.Error("warn msg not found ")
	}
	l.Close()
}
