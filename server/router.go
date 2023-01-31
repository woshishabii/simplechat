package server

import (
	"github.com/CAP/simplechat/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, gin.H{
			"Message": "Hello!",
		})
	})

	r.GET("/ping", func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, gin.H{
			"Message": "Server Started!",
		})
	})

	v1 := r.Group("/api/v1")
	{
		v1.POST("user", api.UserLogin)
		v1.DELETE("user", api.UserLogout)

		v1.POST("room", api.RoomCreate)

		v1.GET("message", api.GetMessage)
		v1.POST("message", api.SendMessage)
	}

	return r
}
