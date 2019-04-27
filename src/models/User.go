package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Username string `gorm:"type:varchar(20);PRIMARY_KEY"`
	Password string `gorm: "type:varchar(15)"`
}

type HeaderHistory struct {
	gorm.Model
	Header string
}

type History struct {
	gorm.Model
	Title   string
	Contect string
}
