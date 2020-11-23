package app

import (
	"github.com/gin-gonic/gin"
)

var ( 
	router 									   = gin.Default() 
)

func StartApplication(){
	// router := gin.Default()
	mapUrlsToken()
	router.Run(":8082") 
}