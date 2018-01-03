package main

import (
	"github.com/labstack/echo"

	"github.com/nuts300/test-echo/controllers"
	"github.com/nuts300/test-echo/db"
)

func main() {
	db := db.GetDB()
	defer db.Close()

	e := echo.New()

	helloController := controllers.NewHelloController()

	e.GET("/hello", helloController.Hello)

	userController := controllers.NewUserController(db)

	e.POST("/users", userController.CreateUser)
	e.GET("/users/:id", userController.GetUser)
	e.GET("/users", userController.GetUsers)
	e.PUT("/users/:id", userController.UpdateUser)
	e.DELETE("/users/:id", userController.DeleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}
