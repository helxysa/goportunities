package main

import (
	 "fmt"
	 "github.com/helxysa/goportunities.git/router"
	 "github.com/helxysa/goportunities.git/config"
)

func main () {

	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return 
	}

	router.Initialize()
	
}


