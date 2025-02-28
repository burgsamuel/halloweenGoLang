package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Spot struct {
	Id        int64    `json:"id"`
	TimeStamp *float64 `json:"time_stamp"`
	Lat       float64  `json:"lat"`
	Lon       float64  `json:"lon"`
	IconUrl   string   `json:"iconUrl"`
}

func main() {

	r := gin.Default()

	r.Static("/static", "./static")

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", nil)
	})

	r.GET("/location", func(ctx *gin.Context) {
		ctx.HTML(200, "addLocation.html", nil)
	})

	r.GET("/mapView", func(ctx *gin.Context) {
		ctx.HTML(200, "mapview.html", nil)
	})

	r.POST("/locationData", createSpot)

	r.GET("/mapData", retrieveSpots)

	r.Run(":8000")
}

func retrieveSpots(c *gin.Context) {

	spots, err := RetrieveMarkers()
	if err != nil {
		c.JSON(500, gin.H{
			"error":   "Failed to retrieve markers",
			"details": err.Error(),
		})
		return
	}

	c.JSON(200, spots)
}

func createSpot(c *gin.Context) {

	var spot Spot

	if err := c.ShouldBindJSON(&spot); err != nil {
		c.JSON(400, gin.H{
			"error":   "Invalid JSON",
			"details": err.Error(),
		})
		fmt.Printf("Binding Error: %v\n", err)
		return
	}

	go InsertSpot(spot)

	c.JSON(200, gin.H{"Message": "Spot Received Successfully"})
}
