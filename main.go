package main

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/nuts300/test-echo/controllers"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// e.POST("/users", saveUser)
	e.GET("/users/:id", controllers.GetUser)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}
