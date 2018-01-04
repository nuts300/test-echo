package userResource

import (
	"github.com/jinzhu/gorm"
	userModel "github.com/nuts300/test-echo/models/user_model"
)

type (
	UserResource interface {
		ReadUserByID(userID int) (userModel.User, error)
		ReadUsers() ([]userModel.User, []error)
		CreateUser(user userModel.User) (userModel.User, error)
		UpdateUser(userID int, user userModel.User) (userModel.User, error)
		DeleteUser(userID int) (userModel.User, error)
	}

	userResource struct {
		db *gorm.DB
	}
)

func (u *userResource) ReadUserByID(userID int) (userModel.User, error) {
	user := userModel.New()
	err := u.db.Where("id = ?", userID).First(&user).Error
	return user, err
}

func (u *userResource) ReadUsers() ([]userModel.User, []error) {
	users := []userModel.User{}
	errors := u.db.Find(&users).GetErrors()
	return users, errors
}

func (u *userResource) CreateUser(user userModel.User) (userModel.User, error) {
	err := u.db.Create(&user).Error
	return user, err
}

func (u *userResource) UpdateUser(userID int, changeFields userModel.User) (userModel.User, error) {
	user := userModel.New()
	err := u.db.Model(&user).Where("id = ?", userID).Updates(&changeFields).Error
	return user, err
}

func (u *userResource) DeleteUser(userID int) (userModel.User, error) {
	user := userModel.New()
	err := u.db.Where("id = ?", userID).Delete(&user).Error
	return user, err
}

func New(db *gorm.DB) UserResource {
	return &userResource{db: db}
}
