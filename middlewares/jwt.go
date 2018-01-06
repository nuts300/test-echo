package middlewares

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/nuts300/test-echo/models"
)

func Jwt() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte("my_secret"),
		TokenLookup: "header:token",
		Claims:      &models.Claims{},
		Skipper: func(c echo.Context) bool {
			if c.Request().URL.Path == "/hello" ||
				(c.Request().URL.Path == "/users" && c.Request().Method == http.MethodPost) ||
				(c.Request().URL.Path == "/auth/login" && c.Request().Method == http.MethodPost) {
				return true
			}
			return false
		},
	})
}
