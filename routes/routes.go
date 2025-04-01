package routes

import (
	"net/http"

	logs "mailscheduler/logger"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, log *logs.Logger) {

	api := router.Group("/api")
	{
		api.GET("/health", Health)
	}
}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
