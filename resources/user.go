package resources

import (
	"github.com/jinzhu/gorm"
	"github.com/nuts300/test-echo/models"
)

type (
	UserResource interface {
		ReadUserByID(userID int) (models.User, error)
		ReadUsers() ([]models.User, []error)
		CreateUser(models.User) (models.User, error)
		UpdateUser(int, models.User) (models.User, error)
		DeleteUser(int) (models.User, error)
	}

	userResource struct {
		db *gorm.DB
	}
)

func (u *userResource) ReadUserByID(userID int) (models.User, error) {
	user := models.NewUser()
	err := u.db.Where("id = ?", userID).First(&user).Error
	return user, err
}

func (u *userResource) ReadUsers() ([]models.User, []error) {
	users := []models.User{}
	errors := u.db.Find(&users).GetErrors()
	return users, errors
}

func (u *userResource) CreateUser(user models.User) (models.User, error) {
	err := u.db.Create(&user).Error
	return user, err
}

func (u *userResource) UpdateUser(userID int, changeFields models.User) (models.User, error) {
	user := models.NewUser()
	err := u.db.Model(&user).Where("id = ?", userID).Updates(&changeFields).Error
	return user, err
}

func (u *userResource) DeleteUser(userID int) (models.User, error) {
	user := models.NewUser()
	err := u.db.Where("id = ?", userID).Delete(&user).Error
	return user, err
}

func NewUserResource(db *gorm.DB) UserResource {
	return &userResource{db: db}
}
