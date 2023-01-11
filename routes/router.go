package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("api/")
	api.Use()
	{
		api.GET("index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"massage": "success",
			})
		})
	}
	return router
}
