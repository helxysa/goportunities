package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/helxysa/goportunities.git/schemas"
	"net/http"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}
	
	if err := ctx.BindJSON(&request); err != nil {
		logger.Errorf("error binding request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err, "error binding request")
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("error validating request: %v", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	sendSucess(ctx, http.StatusCreated, "Opening created successfully")
	ctx.JSON(http.StatusCreated, opening)
}


func sendSucess(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message": msg, 
		"errorCode": code,
	})
}
