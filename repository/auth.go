package repository

import (
	"go/go-api/model"

	"golang.org/x/crypto/bcrypt"
)
func CreateUser(user model.User) {
	encryptPassword := encryptPassword(user.Password)
	userCreate := model.User{Username: user.Username, 
		Password: string(encryptPassword), 
		Fullname: user.Fullname,
		Avatar: user.Avatar}

	model.Db.Create(&userCreate)
}

func encryptPassword(password string) string {
	encrypted,_ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(encrypted)
}