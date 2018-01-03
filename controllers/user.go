package controllers

import (
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	"github.com/nuts300/test-echo/resources"
)

type userController struct {
	resource resources.UserResource
}

type UserController interface {
	GetUser(c echo.Context) error
	GetUsers(c echo.Context) error
	CreateUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}

func (u *userController) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// TODO
	}
	result, err := u.resource.ReadUserByID(id)
	if err != nil {
		// TODO
	}
	return c.JSON(http.StatusOK, result)
}

func (u *userController) GetUsers(c echo.Context) error {
	result, errors := u.resource.ReadUsers()
	if errors != nil {
		// TODO
	}
	return c.JSON(http.StatusOK, result)
}

func (u *userController) CreateUser(c echo.Context) error {
	// user := new(resources.User)
	user := resources.User{}
	if err := c.Bind(user); err != nil {
		// TODO
	}
	result, err := u.resource.CreateUser(user)
	if err != nil {
		// TODO
	}
	return c.JSON(http.StatusOK, result)
}

func (u *userController) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// TODO
	}
	user := resources.User{}
	if err := c.Bind(user); err != nil {
		// TODO
	}
	result, err := u.resource.UpdateUser(id, user)
	if err != nil {
		// TODO
	}
	return c.JSON(http.StatusOK, result)
}

func (u *userController) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// TODO
	}
	result, err := u.resource.DeleteUser(id)
	if err != nil {
		// TODO
	}
	return c.JSON(http.StatusOK, result)
}

func NewUserController(db *gorm.DB) UserController {
	return &userController{resource: resources.NewUserResource(db)}
}
