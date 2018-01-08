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

type (
	userController struct {
		resource resources.UserResource
		logger   appLogger.CustomLogger
		validate *validator.Validate
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
		return appError.NewHTTPError(appError.ErrorInvalidUserID, err)
	}
	result, httpError := u.resource.ReadUserByID(id)
	if httpError != nil {
		return httpError
	}
	return c.JSON(http.StatusOK, result)

}

func (u *userController) GetUsers(c echo.Context) error {
	result, httpError := u.resource.ReadUsers()
	if httpError != nil {
		return httpError
	}
	return c.JSON(http.StatusOK, result)
}

func (u *userController) CreateUser(c echo.Context) error {
	user := models.NewUser()
	if err := c.Bind(user); err != nil {
		return appError.NewHTTPError(appError.ErrorInvalidUserPayload, err)
	}
	if err := u.validate.Struct(user); err != nil {
		return appError.NewHTTPError(appError.ErrorInvalidUserPayload, err)
	}
	createdUser, httpError := u.resource.CreateUser(*user)
	if httpError != nil {
		return httpError
	}
	return c.JSON(http.StatusOK, createdUser)
}

func (u *userController) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return appError.NewHTTPError(appError.ErrorInvalidUserPayload, err)
	}
	user := models.NewUser()
	if err := c.Bind(user); err != nil {
		return appError.NewHTTPError(appError.ErrorInvalidUserPayload, err)
	}
	updatedUser, httpError := u.resource.UpdateUser(id, *user)
	if httpError != nil {
		return httpError
	}
	return c.JSON(http.StatusOK, updatedUser)
}

func (u *userController) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return appError.NewHTTPError(appError.ErrorInvalidUserPayload, err)
	}
	_, httpError := u.resource.DeleteUser(id)
	if httpError != nil {
		return httpError
	}
	return c.NoContent(http.StatusNoContent)
}

func NewUserController(resource resources.UserResource) UserController {
	return &userController{
		resource: resource,
		logger:   appLogger.GetLogger(),
		validate: validator.New(),
	}
}
