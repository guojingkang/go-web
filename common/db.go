package common

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	config := GetConfig()
	db1, err := gorm.Open(mysql.Open(fmt.Sprintf("%v:%v@tcp(127.0.0.1:%v)/go_data?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.User, config.Database.Password, config.Database.Port)), &gorm.Config{})
	db1 = db1.Debug()
	if err != nil {
		panic("failed to connect database")
	}

	db = db1
}

func GetDB() *gorm.DB {
	return db
}
