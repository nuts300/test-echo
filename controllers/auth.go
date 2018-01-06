package controllers

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/nuts300/test-echo/models"
	"github.com/nuts300/test-echo/resources"

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
		WhoAmI(echo.Context) error
	}
)

func (a *authController) Login(c echo.Context) error {
	user := models.NewUser()
	if err := c.Bind(&user); err != nil {
		return appError.NewErrorInvalidUserPayload(err)
	}

	findedUser, aError := a.resource.FindUserByEmailAndPassword(user.Email, user.Password)
	if aError != nil {
		return aError
	}

	token, tokenErr := a.createToken(&findedUser)
	if tokenErr != nil || token == "" {
		return appError.NewErrorUnAuthorized(tokenErr)
	}

	return c.JSON(http.StatusOK, models.AuthInfo{
		ID:    findedUser.ID,
		Email: findedUser.Email,
		Token: token,
	})
}

func (a *authController) RefreshToken(c echo.Context) error {
	token := a.getTokenFromContext(c)
	climes := a.decodeToken(token)
	user := models.User{
		ID:    climes.ID,
		Email: climes.Email,
	}
	newToken, tokenErr := a.createToken(&user)
	if tokenErr != nil || newToken == "" {
		return appError.NewErrorUnAuthorized(tokenErr)
	}

	return c.JSON(http.StatusOK, models.AuthInfo{
		ID:    user.ID,
		Email: user.Email,
		Token: newToken,
	})
}

func (a *authController) WhoAmI(c echo.Context) error {
	token := a.getTokenFromContext(c)
	climes := a.decodeToken(token)
	return c.JSON(http.StatusOK, models.AuthInfo{
		ID:    climes.ID,
		Email: climes.Email,
		Token: token.Raw,
	})
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

func (a *authController) decodeToken(token *jwt.Token) *models.Claims {
	return token.Claims.(*models.Claims)
}

func (a *authController) getTokenFromContext(c echo.Context) *jwt.Token {
	return c.Get("user").(*jwt.Token)
}

func NewAuthController(resource resources.UserResource) AuthController {
	return &authController{resource: resource}
}
