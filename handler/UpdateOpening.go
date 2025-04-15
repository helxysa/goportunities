package handler

import (
	
	"github.com/helxysa/goportunities.git/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
	"errors"
)	



func UpdateOpeningHandler (ctx *gin.Context){
	request := UpdateOpeningRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errorf("error validating request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err, "error validating request")
		return
	}

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

	//Update
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote {
		opening.Remote = request.Remote
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	//save
	if err := db.Save(&opening).Error;err != nil {
		logger.Errorf("error saving opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err, "error saving opening")
		return
	}
	sendSucess(ctx, http.StatusOK, "opening updated successfully")
}

