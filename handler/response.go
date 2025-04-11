package handler 

import (
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, err error) {
	ctx.JSON(code, gin.H{
		"error": err.Error(),
		"errorCode": code, 
	})
}