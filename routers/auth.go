package routers

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/nuts300/test-echo/controllers"
)

func RegisterAuthRoutes(g *echo.Group, db *gorm.DB) {
	authController := controllers.NewAuthController(db)

	g.POST("/login", authController.Login)
	g.POST("/refresh", authController.RefreshToken)
	g.GET("/whoAmi", authController.WhoAmI)
}
