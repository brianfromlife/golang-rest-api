package logger

import (
	"fmt"
)

type Logger struct {
	secret string
}

func (l Logger) Error(message string, err error) {
	fmt.Println(message, err.Error())
}

func (l Logger) Info(message string) {
	fmt.Println(message)
}
