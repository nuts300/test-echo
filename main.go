package main

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/nuts300/test-echo/controllers"
	"github.com/nuts300/test-echo/db"
)

func main() {
	db := db.GetDB()
	defer db.Close()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	userController := controllers.NewUserController(db)

	e.POST("/users", userController.CreateUser)
	e.GET("/users/:id", userController.GetUser)
	e.GET("/users", userController.GetUsers)
	e.PUT("/users/:id", userController.UpdateUser)
	e.DELETE("/users/:id", userController.DeleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}
