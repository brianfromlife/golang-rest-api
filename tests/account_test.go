package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/brianfromlife/golang-ecs/api"
	"github.com/brianfromlife/golang-ecs/api/errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccountSuccess(t *testing.T) {

	// Setup
	requestBody := `{"username":"user","password":"password"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	a := api.App{}

	// Act
	a.RegisterAccount(c)

	// Assert
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, "register account", rec.Body.String())

}

func TestCreateAccountBindError(t *testing.T) {

	// Setup
	requestBody := `{`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	a := api.App{}

	// Act
	a.RegisterAccount(c)

	var error errors.ApiError
	json.Unmarshal([]byte(rec.Body.String()), &error)

	// Assert
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, error.Code, 400)
	assert.Equal(t, error.Name, "BIND_ERROR")

}

func TestCreateAccountValidation(t *testing.T) {

	// Setup
	requestBody := `{"username":"u", "password":"pass"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	a := api.App{}

	// Act
	a.RegisterAccount(c)

	var error errors.ApiError
	json.Unmarshal([]byte(rec.Body.String()), &error)

	// Assert
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, error.Code, 400)
	assert.Equal(t, len(error.Validation), 2)
	assert.Equal(t, error.Name, "VALIDATION")

}
