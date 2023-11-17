package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Logger struct {
	debugFile *os.File
	errorFile *os.File
}

type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Level     string    `json:"level"`
}

func NewLogger(logDir string) (*Logger, error) {
	debugDir := filepath.Join(logDir, "debug")
	errorDir := filepath.Join(logDir, "error")

	if err := os.MkdirAll(debugDir, os.ModePerm); err != nil {
		return nil, err
	}

	if err := os.MkdirAll(errorDir, os.ModePerm); err != nil {
		return nil, err
	}

	debugFile, err := createLogFile(debugDir, "debug")
	if err != nil {
		return nil, err
	}

	errorFile, err := createLogFile(errorDir, "error")
	if err != nil {
		debugFile.Close()
		return nil, err
	}

	return &Logger{
		debugFile: debugFile,
		errorFile: errorFile,
	}, nil
}

func createLogFile(logDir, logType string) (*os.File, error) {
	today := time.Now().Format("2006-01-02")
	logFileName := fmt.Sprintf("%s.log", today)
	logFilePath := filepath.Join(logDir, logFileName)

	file, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (l *Logger) logToFile(file *os.File, level, message string) {
	entry := LogEntry{
		Timestamp: time.Now(),
		Message:   message,
		Level:     level,
	}

	logJSON, err := json.Marshal(entry)
	if err != nil {
		fmt.Println("Error marshaling log entry:", err)
		return
	}

	logJSON = append(logJSON, '\n')

	_, err = file.Write(logJSON)
	if err != nil {
		fmt.Println("Error writing log entry:", err)
	}
}

func (l *Logger) Debug(message string) {
	l.logToFile(l.debugFile, "DEBUG", message)
}

func (l *Logger) Error(message string) {
	l.logToFile(l.errorFile, "ERROR", message)
}

func (l *Logger) Info(message string) {
	fmt.Println(message)
}

func (l *Logger) Close() {
	if l.debugFile != nil {
		l.debugFile.Close()
	}
	if l.errorFile != nil {
		l.errorFile.Close()
	}
}

func main() {
	logger, err := NewLogger("/path/to/logs")
	if err != nil {
		fmt.Println("Error creating logger:", err)
		return
	}
	defer logger.Close()

	// Log messages using the logger like this
	// logger.Debug("This is a debug message.")
	// logger.Error("This is an error message.")
}
