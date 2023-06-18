package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name"`
	Password string `json:"password"`
	Status int `json:"status" gorm:"default:1"` 
}