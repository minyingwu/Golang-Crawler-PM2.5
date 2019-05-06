package routes

import (
	"Go/pm2.5/controllers"

	"github.com/gin-gonic/gin"
)

var ch = make(chan bool)
var isServerOn = false

// StartServer at first time
func StartServer() {
	router := gin.Default()
	router.GET("/", controllers.RespondAllPollutionFromDB)
	router.GET("/:country", controllers.RespondPollutionFromDB)
	go getAirPollutionData()
	if <-ch {
		isServerOn = true
		router.Run(":8888")
	}
}
