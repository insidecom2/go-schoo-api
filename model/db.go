package model

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func InitModel() {
	dsn := os.Getenv("MYSQL_ENDPOINT")
	Db,err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Fail to connect Database")
	}
	Db.AutoMigrate(&User{})
}