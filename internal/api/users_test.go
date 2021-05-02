package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/brianfromlife/golang-ecs/pkg/config"
	usermocks "github.com/brianfromlife/golang-ecs/pkg/mocks/data/users"
	"github.com/brianfromlife/golang-ecs/pkg/services"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRegisterAccount_UsernameExists(t *testing.T) {

	// test data
	testUser := `{"username":"user1","email":"password"}`

	// echo setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(testUser))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// internal setup
	cfg := &config.Settings{}
	mockUserData := usermocks.NewUserDataStoreMock()
	userSvc := services.NewUserService(cfg, mockUserData)

	mockApp := App{
		userSvc: userSvc,
	}

	// act
	mockApp.Register(c)

	// assert
	assert.Equal(t, http.StatusBadRequest, rec.Code)

}
