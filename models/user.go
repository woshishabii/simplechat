package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       string `json:"ID" binding:"required,min=3,max=30"`
	Password string `json:"password" binding:"required,len=64"`
}

type UserLogoutService struct {
	ID   string `json:"ID" bindings:"required"`
	UUID string `json:"UUID" bindings:"required,len=36"`
}
