package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(100);not null"`
	Email    string `json:"email" gorm:"type:varchar(100);unique;unique_index;not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
}

type Users []User
