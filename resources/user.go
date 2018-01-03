package resources

import (
	"github.com/jinzhu/gorm"
)

type (
	User struct {
		Email    string
		Password string
		gorm.Model
	}

	UserResource interface {
		ReadUserByID(userID int) (User, error)
		ReadUsers() ([]User, []error)
		CreateUser(user User) (User, error)
		UpdateUser(userID int, user User) (User, error)
		DeleteUser(userID int) (User, error)
	}

	userResource struct {
		db *gorm.DB
	}
)

func (u *userResource) ReadUserByID(userID int) (User, error) {
	user := User{}
	err := u.db.Where("id = ?", userID).First(&user).Error
	return user, err
}

func (u *userResource) ReadUsers() ([]User, []error) {
	users := []User{}
	errors := u.db.Find(&users).GetErrors()
	return users, errors
}

func (u *userResource) CreateUser(user User) (User, error) {
	err := u.db.Create(&user).Error
	return user, err
}

func (u *userResource) UpdateUser(userID int, changeFields User) (User, error) {
	user := User{}
	err := u.db.Model(&user).Where("id = ?", userID).Updates(&changeFields).Error
	return user, err
}

func (u *userResource) DeleteUser(userID int) (User, error) {
	user := User{}
	err := u.db.Where("id = ?", userID).Delete(&user).Error
	return user, err
}

func NewUserResource(db *gorm.DB) UserResource {
	return &userResource{db: db}
}
