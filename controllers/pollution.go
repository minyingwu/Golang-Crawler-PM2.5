package controllers

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"Go/pm2.5/models"

	"github.com/gin-gonic/gin"
)

// RespondAllPollutionFromDB will get pollution data from DB
func RespondAllPollutionFromDB(c *gin.Context) {
	var p models.Pollution
	// Query all data from DB
	p.FindAllPollution()

	// Show response to client
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.JSON(http.StatusOK, gin.H{"message": p.DataSlice, "status": http.StatusOK})
	fmt.Println("Response finish")
}

// RespondPollutionFromDB will get pollution data from DB
func RespondPollutionFromDB(c *gin.Context) {
	var p models.Pollution

	// Get parameter "country"
	country := c.Param("country")
	// Query only taipei data from DB
	p.FindPollution(bson.M{"county": country})

	// Show response to client
	c.JSON(http.StatusOK, gin.H{"message": p.DataSlice, "status": http.StatusOK})
	fmt.Println("Response finish")
}
