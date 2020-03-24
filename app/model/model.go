package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email  string
	FirstName string
	LastName string
	Password string
	Token string
}
