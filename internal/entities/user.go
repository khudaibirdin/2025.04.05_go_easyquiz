package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Login    string
	Password string
}
