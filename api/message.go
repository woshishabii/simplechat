package api

import (
	"github.com/CAP/simplechat/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMessage(context *gin.Context) {
	var service models.RoomService
	var roomDB models.Room
	err := context.ShouldBindJSON(&service)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{
			"Message": "Bad Request",
		})
		return
	}
	err = models.DB.Where("name = ?", service.Name).First(&roomDB).Error
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{
			"Message": "Room Not Found",
		})
		return
	}
	var messages []models.Message
	models.DB.Where("name = ?", service.Name).Debug().Find(&messages)
	context.IndentedJSON(http.StatusOK, gin.H{
		"Message": "OK GetMessage",
		"Data":    messages,
	})
}

func SendMessage(context *gin.Context) {
	var service models.MessageService
	var roomDB models.Room
	var tokenDB models.Token
	var messageDB models.Message
	err := context.ShouldBindJSON(&service)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{
			"Message": "Bad Request",
		})
		return
	}
	// 检测房间是否存在
	if err := models.DB.Where("name = ?", service.Room).First(&roomDB).Error; err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{
			"Message": "Room Not Found",
		})
		return
	}
	// 检测 Token
	err = models.DB.Where("uuid = ?", service.Token).First(&tokenDB).Error
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{
			"Message": "Token Not Found",
		})
		return
	}
	// 发送
	messageDB.User = tokenDB.ID
	messageDB.Context = service.Context
	messageDB.Name = service.Room
	models.DB.Create(&messageDB)
	context.IndentedJSON(http.StatusOK, gin.H{
		"Message": "OK",
	})
}
