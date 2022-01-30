package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:string(100);not null"`
	Email    string `json:"email" gorm:"type:string(100);unique:unique_index;not null"`
	Password string `json:"password" gorm:"type:string(255);not null"`
}

type LoginCredentials struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
