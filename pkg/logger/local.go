package logger

import "fmt"

type LocalLogger struct{}

func NewLocal() ILogger {
	return LocalLogger{}
}

func (l LocalLogger) Error(message string, err error) {
	fmt.Println(message, err.Error())
}

func (l LocalLogger) Info(message string) {
	fmt.Println(message)
}
