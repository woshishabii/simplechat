package models

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	ID   int    `gorm:"-;primary_key;AUTO_INCREMENT" json:"id"`
	Name string `json:"Name" binding:"required"`
}

type RoomService struct {
	Name string `json:"Name" binding:"required"`
}
