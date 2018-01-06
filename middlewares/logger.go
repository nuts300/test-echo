package middlewares

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Logger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339_nano} [rid]${id} [ip]${remote_ip} [host]${host} [method]${method} [uri]${uri} [status]${status}` + "\n",
		Output: os.Stdout,
	})
}
