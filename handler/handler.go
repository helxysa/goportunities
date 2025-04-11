package handler

import (
	"github.com/helxysa/goportunities.git/config"
	"gorm.io/gorm"
)	


var (
	logger *config.Logger
	db *gorm.DB
)	

func InitializeHandler() {
	logger = config.GetLogger(">> HANDLER: ")
	db = config.GetDB()
}