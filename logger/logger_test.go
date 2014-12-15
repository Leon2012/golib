package logger

import (
	_ "fmt"
	"os"
	"testing"
)

func TestNewLogger(t *testing.T) {
	logger, err := New("test", 2)
	if err != nil {
		t.Fatal(err)
	} else {
		logger.Debug("debug")
		logger.Error("error")
		logger.Info("info")
		logger.Warning("warning")
		logger.Notice("notice")
	}

}

func TestNewFileWorker(t *testing.T) {
	logFile := "./test.log"
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		t.Fatal(err)
	} else {
		worker := NewFileWorker("test", 2, 0, f)
		worker.Minion.Println("hello")
	}
}

func TestNewDailyLogger(t *testing.T) {
	logPath := "./"
	logger, err := NewDailyLogger("App", 0, logPath)
	if err != nil {
		t.Fatal(err)
	} else {
		logger.Debug("debug")
		logger.Error("error")
		logger.Info("info")
		logger.Warning("warning")
		logger.Notice("notice")
	}
}
