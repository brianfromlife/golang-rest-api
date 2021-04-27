package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brianfromlife/golang-ecs/api"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealhCheck(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	a := api.App{}

	a.HealthCheck(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "healthy", rec.Body.String())
}
