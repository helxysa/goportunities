package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/helxysa/goportunities.git/schemas"
	"net/http"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		logger.Errorf("error listing openings: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err, "error listing openings")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "openings listed successfully",
		"data":    openings,
	})
}
