package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	userResource "github.com/nuts300/test-echo/resources/user_resource"
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
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&userResource.User{})

	return db
}
