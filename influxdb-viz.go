package main

import (
	"github.com/gin-gonic/gin"
	"github.com/influxdb/influxdb-viz/controllers"
)

const (
	STATIC           = "static/"
	STATIC_JS        = STATIC + "js"
	STATIC_CSS       = STATIC + "css"
	STATIC_IMAGES    = STATIC + "images"
	STATIC_JSON      = STATIC + "json"
)

func main() {

	router := gin.Default()

	// create routes for all static files
	router.Static("/images", STATIC_IMAGES)
	router.Static("/js", STATIC_JS)
	router.Static("/css", STATIC_CSS)
	router.Static("/json", STATIC_JSON)

	// index page
	homepageController := &controllers.HomepageController{}
	homepageController.Run(router)

	// Add routes for bubbles
	bubbleController := &controllers.BubbleController{}
	bubbleController.Run(router)

	router.Run(":8080")
}
