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
		return appError.NewAppError(appError.INVALID_USER_ID, err)
	}
	if result, err := u.resource.ReadUserByID(id); err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func (u *userController) GetUsers(c echo.Context) error {
	result, err := u.resource.ReadUsers()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}

func (u *userController) CreateUser(c echo.Context) error {
	user := models.NewUser()
	if err := c.Bind(&user); err != nil {
		return appError.NewAppError(appError.INVALID_USER_PAYLOAD, err)
	}
	if err := validate.Struct(user); err != nil {
		return appError.NewAppError(appError.INVALID_USER_PAYLOAD, err)
	}
	if result, err := u.resource.CreateUser(user); err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func (u *userController) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return appError.NewAppError(appError.INVALID_USER_ID, err)
	}
	user := models.NewUser()
	if err := c.Bind(&user); err != nil {
		return appError.NewAppError(appError.INVALID_USER_PAYLOAD, err)
	}
	if result, err := u.resource.UpdateUser(id, user); err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func (u *userController) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return appError.NewAppError(appError.INVALID_USER_ID, err)
	}
	if _, err := u.resource.DeleteUser(id); err != nil {
		return err
	} else {
		return c.NoContent(http.StatusNoContent)
	}
}

func NewUserController(resource resources.UserResource) UserController {
	return &userController{resource: resource}
}
