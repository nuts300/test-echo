package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	helloController "github.com/nuts300/test-echo/controllers/hello_controller"
	userController "github.com/nuts300/test-echo/controllers/user_controller"
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
