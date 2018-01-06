package routers

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/nuts300/test-echo/controllers"
	"github.com/nuts300/test-echo/resources"
)

func RegisterUserRoutes(g *echo.Group, db *gorm.DB) {
	userResource := resources.NewUserResource(db)
	userController := controllers.NewUserController(userResource)

	g.POST("", userController.CreateUser)
	g.GET("/:id", userController.GetUser)
	g.GET("", userController.GetUsers)
	g.PUT("/:id", userController.UpdateUser)
	g.DELETE("/:id", userController.DeleteUser)
}
