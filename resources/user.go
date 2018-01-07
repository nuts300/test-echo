package resources

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/nuts300/test-echo/app_error"
	"github.com/nuts300/test-echo/models"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserResource interface {
		ReadUserByID(userID int) (models.User, *echo.HTTPError)
		ReadUsers() ([]models.User, *echo.HTTPError)
		CreateUser(models.User) (models.User, *echo.HTTPError)
		UpdateUser(int, models.User) (models.User, *echo.HTTPError)
		DeleteUser(int) (models.User, *echo.HTTPError)
		FindUserByEmailAndPassword(email string, password string) (models.User, *echo.HTTPError)
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
			httpError = appError.NewAppError(appError.ErrorNotFoundUser, err)
		default:
			httpError = appError.NewAppError(appError.ErrorFailedReadUser, err)
		}
	}
	return user, httpError
}

func (u *userResource) ReadUsers() ([]models.User, *echo.HTTPError) {
	users := []models.User{}
	errors := u.db.Find(&users).GetErrors()
	var httpError *echo.HTTPError
	if errors != nil {
		httpError = appError.NewAppError(appError.ErrorFailedReadUsers, errors[0])
	}
	return users, httpError
}

func (u *userResource) CreateUser(user models.User) (models.User, *echo.HTTPError) {
	err := u.db.Create(&user).Error
	var httpError *echo.HTTPError
	if err != nil {
		httpError = appError.NewAppError(appError.ErrorFailedCreateUser, err)
	}
	return user, httpError
}

func (u *userResource) UpdateUser(userID int, changeFields models.User) (models.User, *echo.HTTPError) {
	user := models.NewUser()
	var httpError *echo.HTTPError

	result := u.db.Model(&user).Where("id = ?", userID).Updates(&changeFields) // .Error
	if result.Error != nil {
		httpError = appError.NewAppError(appError.ErrorFailedCreateUser, result.Error)
	}

	if err := u.db.Where("id = ?", userID).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			httpError = appError.NewAppError(appError.ErrorNotFoundUser, err)
		default:
			httpError = appError.NewAppError(appError.ErrorFailedReadUser, err)
		}
	}
	return user, httpError
}

func (u *userResource) DeleteUser(userID int) (models.User, *echo.HTTPError) {
	user := models.NewUser()
	var httpError *echo.HTTPError

	result := u.db.Where("id = ?", userID).Delete(&user)

	if result.Error != nil {
		httpError = appError.NewAppError(appError.ErrorFailedDeleteUser, result.Error)
	}
	if result.RowsAffected < 1 {
		httpError = appError.NewAppError(appError.ErrorNotFoundUser, errors.New("Not found user"))
	}
	return user, httpError
}

func (u *userResource) FindUserByEmailAndPassword(email string, password string) (models.User, *echo.HTTPError) {
	user := models.NewUser()
	err := u.db.Where("email = ?", email).First(&user).Error
	var httpError *echo.HTTPError
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			httpError = appError.NewAppError(appError.ErrorNotFoundUser, err)
		default:
			httpError = appError.NewAppError(appError.ErrorFailedReadUser, err)
		}
		return user, httpError
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		httpError = appError.NewAppError(appError.ErrorUnauthorized, err)
	}

	return user, httpError
}

func NewUserResource(db *gorm.DB) UserResource {
	return &userResource{db: db}
}
