package handlers

import (
	"net/http"

	"github.com/brianfromlife/golang-ecs/pkg/logger"
	"github.com/brianfromlife/golang-ecs/pkg/models"
	"github.com/brianfromlife/golang-ecs/pkg/services"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	logger  logger.ILogger
	userSvc services.IUserService
}

func NewUserHandler(logger logger.ILogger, userSvc services.IUserService) *UserHandler {
	return &UserHandler{
		logger:  logger,
		userSvc: userSvc,
	}
}

func (h UserHandler) Register(c echo.Context) error {
	newUser, err := models.ValidateRegisterRequest(c)
	if err != nil {
		h.logger.Error("Error validating register request", err.Error)
		return c.JSON(err.Code, err)
	}
	err = h.userSvc.CreateAccount(newUser)

	if err != nil {
		h.logger.Error("Error validating register request", err.Error)
		return c.JSON(err.Code, err)
	}

	return c.String(http.StatusOK, "register")
}
