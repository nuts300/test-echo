package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/nuts300/test-echo/app_error"
	"github.com/nuts300/test-echo/app_logger"
	"github.com/nuts300/test-echo/models"
	"github.com/nuts300/test-echo/resources"
)

var logger = appLogger.GetLogger()
var validate = validator.New()

type (
	userController struct {
		resource resources.UserResource
	}

	UserController interface {
		GetUser(echo.Context) error
		GetUsers(echo.Context) error
		CreateUser(echo.Context) error
		UpdateUser(echo.Context) error
		DeleteUser(echo.Context) error
	}
)

func (u *userController) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return appError.ErrorInvalidUserID(err)
	}
	result, aError := u.resource.ReadUserByID(id)
	if aError != nil {
		return aError
	}

	return c.JSON(http.StatusOK, result)
}

func (u *userController) GetUsers(c echo.Context) error {
	result, aError := u.resource.ReadUsers()
	if aError != nil {
		return aError
	}
	return c.JSON(http.StatusOK, result)
}

func (u *userController) CreateUser(c echo.Context) error {
	user := models.NewUser()
	if err := c.Bind(&user); err != nil {
		return appError.ErrorInvalidUserPayload(err)
	}
	if err := validate.Struct(user); err != nil {
		return appError.ErrorInvalidUserID(err)
	}
	result, aError := u.resource.CreateUser(user)
	if aError != nil {
		return aError
	}

	return c.JSON(http.StatusOK, result)
}

func (u *userController) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return appError.ErrorInvalidUserID(err)
	}
	user := models.NewUser()
	if err := c.Bind(&user); err != nil {
		return appError.ErrorInvalidUserPayload(err)
	}
	result, aError := u.resource.UpdateUser(id, user)
	if aError != nil {
		return aError
	}

	return c.JSON(http.StatusOK, result)
}

func (u *userController) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return appError.ErrorInvalidUserID(err)
	}
	_, aError := u.resource.DeleteUser(id)
	if aError != nil {
		return aError
	}

	return c.NoContent(http.StatusNoContent)
}

func NewUserController(resource resources.UserResource) UserController {
	return &userController{resource: resource}
}
