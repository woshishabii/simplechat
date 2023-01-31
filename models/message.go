package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	ID      int    `gorm:"-;primary_key;AUTO_INCREMENT" json:"id"`
	User    string `json:"User" binding:"required"`
	Context string `json:"Context" binding:"required"`
	Room
}

type MessageService struct {
	Token   string `json:"Token" binding:"required,len=36"`
	Context string `json:"Context" binding:"required"`
	Room    string `json:"Room" binding:"required"`
}
