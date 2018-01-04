package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	helloController "github.com/nuts300/test-echo/controllers/hello_controller"
	userController "github.com/nuts300/test-echo/controllers/user_controller"
	"github.com/nuts300/test-echo/db"
)

func main() {
	db := db.GetDB()
	defer db.Close()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	hello := helloController.New()

	e.GET("/hello", hello.Hello)

	user := userController.New(db)

	e.POST("/users", user.CreateUser)
	e.GET("/users/:id", user.GetUser)
	e.GET("/users", user.GetUsers)
	e.PUT("/users/:id", user.UpdateUser)
	e.DELETE("/users/:id", user.DeleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}
