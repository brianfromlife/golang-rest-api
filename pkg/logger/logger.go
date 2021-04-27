package logger

import "fmt"

type ILogger interface {
	Error(message string, err error)
}

type Logger struct{}

func NewLocal() ILogger {
	return Logger{}
}

func (l Logger) Error(message string, err error) {
	fmt.Println(message, err.Error())
}
