package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func InitModel() {
	dsn := "root:root@tcp(127.0.0.1:3306)/go_api?charset=utf8mb4&parseTime=True&loc=Local"
	Db,err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Fail to connect Database")
	}
	Db.AutoMigrate(&User{})
}