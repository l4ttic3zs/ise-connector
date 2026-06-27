// Package logger is used to log in the app
package logger

import (
	"fmt"
	"os"
	"time"
)

type LogLevel string

const (
	InfoLevel    LogLevel = "INFO"
	ErrorLevel   LogLevel = "ERROR"
	WarningLevel LogLevel = "Warning"
)

type LogMessage struct {
	Timestamp time.Time
	Level     LogLevel
	Message   string
}

type AsyncLogger struct {
	logChan  chan LogMessage
	stopChan chan struct{}
}

func New(bufferSize int) *AsyncLogger {
	l := &AsyncLogger{
		logChan:  make(chan LogMessage, bufferSize),
		stopChan: make(chan struct{}),
	}

	go l.start()

	return l
}

func (l *AsyncLogger) start() {
	for {
		select {
		case msg := <-l.logChan:
			fmt.Fprintf(os.Stdout, "[%s] %s: %s\n",
				msg.Timestamp.Format("2006-01-02 15:04:05"),
				msg.Level,
				msg.Message,
			)
		case <-l.stopChan:
			close(l.logChan)
			for msg := range l.logChan {
				fmt.Fprintf(os.Stdout, "[%s] %s: %s (FLUSHED)\n", msg.Timestamp.Format("2006-01-02 15:04:05"), msg.Level, msg.Message)
			}
			return
		}
	}
}

func (l *AsyncLogger) Info(msg string) {
	l.logChan <- LogMessage{
		Timestamp: time.Now(),
		Level:     InfoLevel,
		Message:   msg,
	}
}

func (l *AsyncLogger) Error(msg string) {
	l.logChan <- LogMessage{
		Timestamp: time.Now(),
		Level:     ErrorLevel,
		Message:   msg,
	}
}

func (l *AsyncLogger) Close() {
	close(l.stopChan)
}
