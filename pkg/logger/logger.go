package logger

import (
	"fmt"

	"github.com/brianfromlife/golang-ecs/pkg/config"
)

type ILogger interface {
	Error(message string, err error)
	Info(message string)
}

type Logger struct {
	secret string
}

func NewLogger(cfg *config.Settings) ILogger {
	if cfg.Env == "development" {
		return NewLocal()
	}
	return &Logger{secret: "secret"}
}

func (l Logger) Error(message string, err error) {
	fmt.Println(message, err.Error())
}

func (l Logger) Info(message string) {
	fmt.Println(message)
}
