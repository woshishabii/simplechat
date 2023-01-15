package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"os"
)

type User struct {
	gorm.Model
	Email    string
	Password string
	Nickname string
	Sex      bool
}

func main() {
	// 获取环境配置
	err := godotenv.Load()
	if err != nil {
		panic("Error while loading environment configuration")
	}

	// 连接数据库
	db, err := gorm.Open(sqlite.Open("simple_chat.db"), &gorm.Config{})
	if err != nil {
		panic("Error while connecting to database")
	}

	// 迁移数据库结构
	db.AutoMigrate(&User{})

	// 初始化GIN
	// gin.SetMode("release")
	r := gin.Default()

	// 初始化接口
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Message": "pong",
		})
	})

	r.POST("/v1/user", createUser)

	// 运行
	r.Run(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")))
}

func createUser(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	// fmt.Printf("%v", newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}
