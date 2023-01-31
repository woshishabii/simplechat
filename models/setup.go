package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("simple_chat.db"), &gorm.Config{})

	if err != nil {
		panic("Error while connecting to database")
	}

	err = database.AutoMigrate(&User{}, &Token{}, &Room{}, &Message{})

	if err != nil {
		panic("Error while migrating models")
	}

	DB = database
}
