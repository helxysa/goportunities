package router

import (
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine){
	v1 := router.Group("/api/v1/")
	{
		v1.GET("/opening", func( *gin.Context) {
			v1.JSON(200, gin.H{
				"message": "pong",
			  })
		})
		v1.POST("/opening", createOpening)
	}
}
