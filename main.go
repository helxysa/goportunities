package main

import (

	 "github.com/helxysa/goportunities.git/router"
	 "github.com/helxysa/goportunities.git/config"
)

var (
	logger *config.Logger
)


func main () {
	logger = config.GetLogger(">> MAIN: ")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return 
	}

	router.Initialize()
	
}


