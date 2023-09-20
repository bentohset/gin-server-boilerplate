package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255; not null; unique" json:"username"`
	Password string `gorm:"size:255; not null; unique" json:"-"`
	Entries  []Entry
}

type UserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"` //validation
}
