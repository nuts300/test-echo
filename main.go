package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/nuts300/test-echo/controllers"
	"github.com/nuts300/test-echo/db"
	errorHander "github.com/nuts300/test-echo/error_handler"
	"github.com/nuts300/test-echo/models"
)

func main() {
	db := db.GetDB()
	defer db.Close()

	e := echo.New()
	e.HTTPErrorHandler = errorHander.AppErrorHandler

	// e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339_nano} [rid]${id} [ip]${remote_ip} [host]${host} [method]${method} [uri]${uri} [status]${status}` + "\n",
		Output: os.Stdout,
	}))

	jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte("my_secret"),
		TokenLookup: "header:token",
		Claims:      &models.Claims{},
		Skipper: func(c echo.Context) bool {
			if c.Request().URL.Path == "/hello" ||
				(c.Request().URL.Path == "/users" && c.Request().Method == http.MethodPost) {
				return true
			}
			return false
		},
	})
	e.Use(jwtMiddleware)

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
