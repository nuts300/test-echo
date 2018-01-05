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

	var httpError *echo.HTTPError
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			httpError = appError.NewAppError(appError.NOT_FOUND_USER, err)
		default:
			httpError = appError.NewAppError(appError.FAILED_READ_USER, err)
		}
	}
	return user, httpError
}

func (u *userResource) ReadUsers() ([]models.User, *echo.HTTPError) {
	users := []models.User{}
	errors := u.db.Find(&users).GetErrors()
	var httpError *echo.HTTPError
	if errors != nil {
		httpError = appError.NewAppError(appError.FAILED_READ_USERS, errors[0])
	}
	return users, httpError
}

func (u *userResource) CreateUser(user models.User) (models.User, *echo.HTTPError) {
	err := u.db.Create(&user).Error
	var httpError *echo.HTTPError
	if err != nil {
		httpError = appError.NewAppError(appError.FAILED_CREATE_USER, err)
	}
	return user, httpError
}

func (u *userResource) UpdateUser(userID int, changeFields models.User) (models.User, *echo.HTTPError) {
	user := models.NewUser()
	var httpError *echo.HTTPError

	err := u.db.Model(&user).Where("id = ?", userID).Updates(&changeFields).Error
	if err != nil {
		httpError = appError.NewAppError(appError.FAILED_UPDATE_USER, err)
	}

	if err := u.db.Where("id = ?", userID).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			httpError = appError.NewAppError(appError.NOT_FOUND_USER, err)
		default:
			httpError = appError.NewAppError(appError.FAILED_READ_USER, err)
		}
	}
	return user, httpError
}

func (u *userResource) DeleteUser(userID int) (models.User, *echo.HTTPError) {
	user := models.NewUser()
	var httpError *echo.HTTPError
	if err := u.db.Where("id = ?", userID).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			httpError = appError.NewAppError(appError.NOT_FOUND_USER, err)
		default:
			httpError = appError.NewAppError(appError.FAILED_READ_USER, err)
		}
		return user, httpError
	}

	err := u.db.Where("id = ?", userID).Delete(&user).Error
	if err != nil {
		httpError = appError.NewAppError(appError.FAILED_DELETE_USER, err)
	}
	return user, httpError
}

func NewUserResource(db *gorm.DB) UserResource {
	return &userResource{db: db}
}
