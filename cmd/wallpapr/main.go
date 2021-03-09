package main

import (
	"log"
	"os"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		client.Logger.Println(err)
		os.Exit(1)
	}
}

func exitOnError(err error) {
	if err != nil {
		client.Logger.Println(err)
		os.Exit(1)
	}
}

// BuiltinLogger implements logger methods
type BuiltinLogger struct {
	logger *log.Logger
}

// NewBuiltinLogger initializes a BuiltinLogger using os.Stdout
func NewBuiltinLogger() *BuiltinLogger {
	return &BuiltinLogger{logger: log.New(os.Stdout, "", 5)}
}

// Print implements Print method of the Logger interface
func (l *BuiltinLogger) Print(v ...interface{}) {
	l.logger.Print(v...)
}

// Printf implements Printf method of the Logger interface
func (l *BuiltinLogger) Printf(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}

// Println implements Println method of the Logger interface
func (l *BuiltinLogger) Println(v ...interface{}) {
	l.logger.Println(v...)
}
