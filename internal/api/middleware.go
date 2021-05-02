package api

import (
	"fmt"

	"github.com/brianfromlife/golang-ecs/pkg/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type Middleware struct {
	config *config.Settings
}

func (m Middleware) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return echo.ErrUnauthorized
		}

		type Claims struct {
			Id  string `json:"id"`
			Exp int    `json:"exp"`
			jwt.StandardClaims
		}

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.config.JwtSecret), nil
		})

		if err != nil {
			fmt.Println(err)
			return echo.ErrUnauthorized
		}

		claims, ok := token.Claims.(*Claims)

		if ok && token.Valid {
			c.Set("user", claims.Id)
			return next(c)
		} else {
			return echo.ErrUnauthorized
		}

	}
}
