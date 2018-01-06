package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/nuts300/test-echo/controllers"
	"github.com/nuts300/test-echo/db"
	errorHander "github.com/nuts300/test-echo/error_handler"
	appMiddleware "github.com/nuts300/test-echo/middlewares"
	"github.com/nuts300/test-echo/routers"
)

func main() {
	db := db.GetDB()
	defer db.Close()

	e := echo.New()
	e.HTTPErrorHandler = errorHander.AppErrorHandler

	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(appMiddleware.Cors())
	e.Use(appMiddleware.Logger())
	e.Use(appMiddleware.Jwt())

	helloController := controllers.NewHelloController()
	e.GET("/hello", helloController.Hello)

	routers.RegisterUserRoutes(e.Group("/users"), db)
	routers.RegisterAuthRoutes(e.Group("/auth"), db)

	e.Logger.Fatal(e.Start(":1323"))
}
