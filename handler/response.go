package handler 

import (
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, err error, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"error": err.Error(),
		"message": msg, 
		"errorCode": code,
	})
}