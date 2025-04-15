package handler

import (
	"fmt"
	"github.com/gin-gonic/gin" 
	"github.com/helxysa/goportunities.git/schemas"
	"net/http"
)	



func ShowOpeningHandler (ctx *gin.Context){
	id := ctx.Query("id")
	if id == ""{
		sendError(ctx, http.StatusBadRequest, fmt.Errorf("id parameter is empty"), "id parameter is empty")
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Errorf("opening not found"), "opening not found")
		return
	}

	sendSucess(ctx, http.StatusOK, "opening found successfully")
	ctx.JSON(http.StatusOK, opening)
}
