package api

import (
	"github.com/CAP/simplechat/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RoomCreate(context *gin.Context) {
	var service models.RoomService
	var roomDB models.Room
	var room models.Room
	err := context.ShouldBindJSON(&service)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{
			"Message": "Bad Request",
		})
		return
	}
	if err := models.DB.Where("name = ?", service.Name).First(&roomDB).Error; err == nil {
		context.IndentedJSON(http.StatusFound, gin.H{
			"Message": "Room Already Exists",
		})
		return
	}
	room.Name = service.Name
	models.DB.Create(&room)
	context.IndentedJSON(http.StatusOK, gin.H{
		"Message": "Room Created",
		"Room":    room,
	})
}
