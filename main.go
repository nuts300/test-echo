package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/nuts300/test-echo/controllers"
	"github.com/nuts300/test-echo/db"
	errorHander "github.com/nuts300/test-echo/error_handler"
	appMiddleware "github.com/nuts300/test-echo/middlewares"
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

	userController := controllers.NewUserController(db)

	e.POST("/users", userController.CreateUser)
	e.GET("/users/:id", userController.GetUser)
	e.GET("/users", userController.GetUsers)
	e.PUT("/users/:id", userController.UpdateUser)
	e.DELETE("/users/:id", userController.DeleteUser)

	authController := controllers.NewAuthController(db)

	e.POST("/auth/login", authController.Login)
	e.POST("/auth/refresh", authController.RefreshToken)
	e.GET("/auth/whoAmi", authController.WhoAmI)

	e.Logger.Fatal(e.Start(":1323"))
}
