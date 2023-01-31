package api

import (
	"github.com/CAP/simplechat/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func UserLogin(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	var token = models.Token{ID: user.ID, UUID: uuid.NewString()}
	var UserDB models.User
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{
			"Message": "Bad Request!",
		})
		return
	}
	if err := models.DB.Where("id = ?", user.ID).First(&UserDB).Error; err != nil {
		models.DB.Create(&user)
		models.DB.Create(&token)
		context.IndentedJSON(http.StatusCreated, gin.H{
			"Message": "Account Created",
			"UUID":    token.UUID,
		})
		return
	}
	if UserDB.Password != user.Password {
		context.IndentedJSON(http.StatusForbidden, gin.H{
			"Message": "Password Wrong!",
		})
		return
	}
	var TokenDB models.Token
	if err := models.DB.Where("id = ?", user.ID).First(&TokenDB).Error; err == nil {
		context.IndentedJSON(http.StatusForbidden, gin.H{
			"Message": "Already Logged In",
		})
		return
	}
	models.DB.Create(&token)
	context.IndentedJSON(http.StatusAccepted, gin.H{
		"Message": "Login Successful",
		"UUID":    token.UUID,
	})
}

func UserLogout(context *gin.Context) {
	var service models.UserLogoutService
	var token models.Token
	err := context.ShouldBindJSON(&service)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{
			"Message": "Bad Request",
		})
		return
	}
	if err := models.DB.Where("id = ?", service.ID).First(&token).Error; err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{
			"Message": "Token Not Found",
		})
		return
	}
	if token.UUID != service.UUID {
		context.IndentedJSON(http.StatusForbidden, gin.H{
			"Message": "Wrong UUID",
		})
		return
	}
	models.DB.Delete(&token)
	context.IndentedJSON(http.StatusOK, gin.H{
		"Message": "Logout Successful",
	})
}
