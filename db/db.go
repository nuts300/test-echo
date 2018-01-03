package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

/*
GetDB get db connection
*/
func GetDB() *gorm.DB {
	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		// TODO connection error handling
		panic(err)
	}
	return db
}
