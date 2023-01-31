package conf

import (
	"github.com/CAP/simplechat/models"
	"github.com/gin-gonic/gin"
)

func Init() {
	gin.SetMode(gin.ReleaseMode)
	models.ConnectDatabase()
}
