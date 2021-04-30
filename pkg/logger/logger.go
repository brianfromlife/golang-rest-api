package logger

import (
	"github.com/brianfromlife/golang-ecs/pkg/config"
)

type ILogger interface {
	Error(message string, err error)
	Info(message string)
}

func NewLogger(cfg *config.Settings) ILogger {
	if cfg.Env == "development" {
		return &LocalLogger{}
	}
	return &Logger{secret: "secret"}
}
