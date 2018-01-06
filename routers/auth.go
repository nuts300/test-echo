package routers

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/nuts300/test-echo/controllers"
	"github.com/nuts300/test-echo/resources"
)

func RegisterAuthRoutes(g *echo.Group, db *gorm.DB) {
	userResource := resources.NewUserResource(db)
	authController := controllers.NewAuthController(userResource)

	g.POST("/login", authController.Login)
	g.POST("/refresh", authController.RefreshToken)
	g.GET("/whoAmi", authController.WhoAmI)
}
