package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/nuts300/test-echo/models"
	"github.com/nuts300/test-echo/resources"

	"github.com/jinzhu/gorm"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/nuts300/test-echo/app_error"
)

type (
	authController struct {
		resource resources.UserResource
	}

	AuthController interface {
		Login(echo.Context) error
		RefreshToken(echo.Context) error
	}
)

func (a *authController) Login(c echo.Context) error {
	user := models.NewUser()
	if err := c.Bind(&user); err != nil {
		return appError.NewAppError(appError.INVALID_USER_PAYLOAD, err)
	}

	findedUser, err := a.resource.FindUserByEmailAndPassword(user.Email, user.Password)
	if err != nil {
		return err
	}
	token, tokenErr := a.createToken(&findedUser)
	if tokenErr != nil || token == "" {
		return appError.NewAppError(appError.UNAUTHORIZED_ERROR, tokenErr)
	}

	return c.JSON(http.StatusOK, token)
}

func (a *authController) RefreshToken(c echo.Context) error {
	return errors.New("Not implement")
}

func (a *authController) createToken(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	exp := time.Now().Add(time.Hour * 24 * 1).Unix()
	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = exp
	tokenString, err := token.SignedString([]byte("my_secret"))
	return tokenString, err
}

func NewAuthController(db *gorm.DB) AuthController {
	return &authController{resource: resources.NewUserResource(db)}
}
