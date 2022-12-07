package model

import "gorm.io/gorm"


type Course struct {
	gorm.Model
	Name  string `gorm:"unique"`
	Status  string
  }