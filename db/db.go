package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/nuts300/test-echo/models"
)

/*
GetDB get db connection
*/
func GetDB() *gorm.DB {
	// db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", "root@/testdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		// TODO connection error handling
		panic(err)
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{})

	return db
}
