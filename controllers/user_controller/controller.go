package userController

import (
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	appError "github.com/nuts300/test-echo/app_error"
	appLogger "github.com/nuts300/test-echo/app_logger"
	userModel "github.com/nuts300/test-echo/models/user_model"
	userResource "github.com/nuts300/test-echo/resources/user_resource"
)

var logger = appLogger.GetLogger()

type (
	userController struct {
		resource userResource.UserResource
	}

	UserController interface {
		GetUser(c echo.Context) error
		GetUsers(c echo.Context) error
		CreateUser(c echo.Context) error
		UpdateUser(c echo.Context) error
		DeleteUser(c echo.Context) error
	}
)

func (u *userController) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// TODO
		// echo.NewHTTPError()
		// return errors.Wrap(err, `Invalide user id ${c.Param("id")}`)
		return appError.NewAppError(appError.INVALID_USER_ID, err)
	}
	result, err := u.resource.ReadUserByID(id)
	if err != nil {
		// TODO
		return err
	}
	return c.JSON(http.StatusOK, result)
}

func (u *userController) GetUsers(c echo.Context) error {
	result, errors := u.resource.ReadUsers()
	if errors != nil {
		// TODO
		return errors[0]
	}
	return c.JSON(http.StatusOK, result)
}

func (u *userController) CreateUser(c echo.Context) error {
	// user := new(resources.User)
	user := userModel.New()
	if err := c.Bind(&user); err != nil {
		// TODO
		return err
	}
	result, err := u.resource.CreateUser(user)
	if err != nil {
		// TODO
		return err
	}
	return c.JSON(http.StatusOK, result)
}

func (u *userController) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// TODO
		return err
	}
	user := userModel.New()
	if err := c.Bind(&user); err != nil {
		// TODO
		return err
	}
	result, err := u.resource.UpdateUser(id, user)
	if err != nil {
		// TODO
		return err
	}
	return c.JSON(http.StatusOK, result)
}

func (u *userController) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// TODO
		return err
	}
	result, err := u.resource.DeleteUser(id)
	if err != nil {
		// TODO
		return err
	}
	return c.JSON(http.StatusOK, result)
}

func New(db *gorm.DB) UserController {
	return &userController{resource: userResource.New(db)}
}
