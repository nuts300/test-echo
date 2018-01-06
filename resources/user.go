package resources

import (
	"github.com/jinzhu/gorm"
	"github.com/nuts300/test-echo/app_error"
	"github.com/nuts300/test-echo/models"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserResource interface {
		ReadUserByID(userID int) (models.User, *appError.AppError)
		ReadUsers() ([]models.User, *appError.AppError)
		CreateUser(models.User) (models.User, *appError.AppError)
		UpdateUser(int, models.User) (models.User, *appError.AppError)
		DeleteUser(int) (models.User, *appError.AppError)
		FindUserByEmailAndPassword(email string, password string) (models.User, *appError.AppError)
	}

	userResource struct {
		db *gorm.DB
	}
)

func (u *userResource) ReadUserByID(userID int) (models.User, *appError.AppError) {
	user := models.NewUser()
	err := u.db.Where("id = ?", userID).First(&user).Error

	var aError *appError.AppError
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			aError = appError.NewErrorNotFoundUser(err)
		default:
			aError = appError.NewErrorFailedReadUser(err)
		}
	}
	return user, aError
}

func (u *userResource) ReadUsers() ([]models.User, *appError.AppError) {
	users := []models.User{}
	errors := u.db.Find(&users).GetErrors()
	var aError *appError.AppError
	if errors != nil {
		aError = appError.NewErrorFailedReadUsers(errors[0])
	}
	return users, aError
}

func (u *userResource) CreateUser(user models.User) (models.User, *appError.AppError) {
	err := u.db.Create(&user).Error
	var aError *appError.AppError
	if err != nil {
		aError = appError.NewErrorFailedCreateUser(err)
	}
	return user, aError
}

func (u *userResource) UpdateUser(userID int, changeFields models.User) (models.User, *appError.AppError) {
	user := models.NewUser()
	var aError *appError.AppError

	err := u.db.Model(&user).Where("id = ?", userID).Updates(&changeFields).Error
	if err != nil {
		aError = appError.NewErrorFailedUpdateUser(err)
	}

	if err := u.db.Where("id = ?", userID).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			aError = appError.NewErrorNotFoundUser(err)
		default:
			aError = appError.NewErrorFailedReadUser(err)
		}
	}
	return user, aError
}

func (u *userResource) DeleteUser(userID int) (models.User, *appError.AppError) {
	user := models.NewUser()
	var aError *appError.AppError

	result := u.db.Where("id = ?", userID).Delete(&user)

	if result.Error != nil {
		aError = appError.NewErrorFailedDeleteUser(result.Error)
	}
	if result.RowsAffected < 1 {
		aError = appError.NewErrorNotFoundUser(result.Error)
	}
	return user, aError
}

func (u *userResource) FindUserByEmailAndPassword(email string, password string) (models.User, *appError.AppError) {
	user := models.NewUser()
	err := u.db.Where("email = ?", email).First(&user).Error
	var aError *appError.AppError
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			aError = appError.NewErrorNotFoundUser(err)
		default:
			aError = appError.NewErrorFailedReadUser(err)
		}
		return user, aError
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		aError = appError.NewErrorNotFoundUser(err)
	}

	return user, aError
}

func NewUserResource(db *gorm.DB) UserResource {
	return &userResource{db: db}
}
