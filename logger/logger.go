package logs

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	log  *logrus.Logger
	file *os.File
	mu   sync.Mutex
}

// NewLogger initializes the logger and returns a Logger instance
func NewLogger(logFilePath string) (*Logger, error) {
	log := logrus.New()
	err := os.MkdirAll("logs/", 0777)
	if err != nil {
		return nil, err
	}
	// Open the log file once
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	log.SetOutput(logFile)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	return &Logger{
		log:  log,
		file: logFile,
	}, nil
}

func (l *Logger) Info(msg string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.log.Info(msg)
}

func (l *Logger) Warn(msg string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.log.Warn(msg)
}

func (l *Logger) Error(msg string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.log.Error(msg)
}

func (l *Logger) Fatal(msg string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.log.Fatal(msg)
}

func (l *Logger) Close() {
	l.mu.Lock()
	defer l.mu.Unlock()
	if err := l.file.Close(); err != nil {
		l.log.Error("Failed to close log file:", err)
	}
}
