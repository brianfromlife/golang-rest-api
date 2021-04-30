package logger

import "fmt"

type ILogger interface {
	Error(message string, err error)
	Info(message string)
}

type Logger struct {
	secret string
}

func NewLogger(secret string) ILogger {
	return Logger{secret: secret}
}

func (l Logger) Error(message string, err error) {
	fmt.Println(message, err.Error())
}

func (l Logger) Info(message string) {
	fmt.Println(message)
}
