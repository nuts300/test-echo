package resources

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/nuts300/test-echo/app_error"
	"github.com/nuts300/test-echo/models"
)

type (
	UserResource interface {
		ReadUserByID(userID int) (models.User, *echo.HTTPError)
		ReadUsers() ([]models.User, *echo.HTTPError)
		CreateUser(models.User) (models.User, *echo.HTTPError)
		UpdateUser(int, models.User) (models.User, *echo.HTTPError)
		DeleteUser(int) (models.User, *echo.HTTPError)
	}

	userResource struct {
		db *gorm.DB
	}
)

func (u *userResource) ReadUserByID(userID int) (models.User, *echo.HTTPError) {
	user := models.NewUser()
	err := u.db.Where("id = ?", userID).First(&user).Error

	var e *echo.HTTPError = nil
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			e = appError.NewAppError(appError.NOT_FOUND_USER, err)
		default:
			e = appError.NewAppError(appError.FAILED_READ_USER, err)
		}
	}
	return user, e
}

func (u *userResource) ReadUsers() ([]models.User, *echo.HTTPError) {
	users := []models.User{}
	errors := u.db.Find(&users).GetErrors()
	var e *echo.HTTPError = nil
	if errors != nil {
		e = appError.NewAppError(appError.FAILED_READ_USERS, errors[0])
	}
	return users, e
}

func (u *userResource) CreateUser(user models.User) (models.User, *echo.HTTPError) {
	err := u.db.Create(&user).Error
	var e *echo.HTTPError = nil
	if err != nil {
		e = appError.NewAppError(appError.FAILED_CREATE_USER, err)
	}
	return user, e
}

func (u *userResource) UpdateUser(userID int, changeFields models.User) (models.User, *echo.HTTPError) {
	user := models.NewUser()
	err := u.db.Model(&user).Where("id = ?", userID).Updates(&changeFields).Error
	var e *echo.HTTPError = nil
	if err != nil {
		e = appError.NewAppError(appError.FAILED_UPDATE_USER, err)
	}
	return user, e
}

func (u *userResource) DeleteUser(userID int) (models.User, *echo.HTTPError) {
	user := models.NewUser()
	err := u.db.Where("id = ?", userID).Delete(&user).Error
	var e *echo.HTTPError = nil
	if err != nil {
		e = appError.NewAppError(appError.FAILED_DELETE_USER, err)
	}
	return user, e
}

func NewUserResource(db *gorm.DB) UserResource {
	return &userResource{db: db}
}
