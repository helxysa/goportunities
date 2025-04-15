package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/helxysa/goportunities.git/schemas"
)	



func DeleteOpeningHandler (ctx *gin.Context){
	id := ctx.Query("id")
	if id == "" {
		logger.Errorf("id is required")
		sendError(ctx, http.StatusBadRequest, errors.New("id is required"), "id parameter is empty")
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf("error finding opening: %v", err.Error())
		sendError(ctx, http.StatusNotFound, err, "opening not found")
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		logger.Errorf("error deleting opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err, "error deleting opening")
		return
	}

	sendSucess(ctx, http.StatusOK, "opening deleted successfully")
}
