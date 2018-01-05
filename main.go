package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/nuts300/test-echo/controllers"
	"github.com/nuts300/test-echo/db"
	errorHander "github.com/nuts300/test-echo/error_handler"
)

func main() {
	db := db.GetDB()
	defer db.Close()

	e := echo.New()
	e.HTTPErrorHandler = errorHander.AppErrorHandler

	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339_nano} [ip]${remote_ip} [host]${host} [method]${method} [uri]${uri} [status]${status}` + "\n",
		Output: os.Stdout}))

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
