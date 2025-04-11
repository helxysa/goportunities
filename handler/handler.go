package handler

import (
	 "github.com/gin-gonic/gin" 
	  "net/http"
)	



func CreateOpeningHandler (ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H {
		"message":"GET",
	})
}

func ShowOpeningHandler (ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H {
		"message":"GET",
	})
}

func UpdateOpeningHandler (ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H {
		"message":"GET",
	})
}


func DeleteOpeningHandler (ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H {
		"message":"GET",
	})
}

func ListOpeningHandler (ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H {
		"message":"GET",
	})
}

