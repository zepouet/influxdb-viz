package main

import (
	"github.com/gin-gonic/gin"
	"github.com/influxdb/influxdb-viz/controllers"
	"github.com/influxdb/influxdb-viz/configuration"
)

const (
	STATIC           = "static/"
	STATIC_JS        = STATIC + "js"
	STATIC_CSS       = STATIC + "css"
	STATIC_IMAGES    = STATIC + "images"
	STATIC_TAGS    	 = STATIC + "tags"
	STATIC_JSON      = STATIC + "json"
)

func main() {

	router := gin.Default()

	// create routes for all static files
	router.Static("/images", STATIC_IMAGES)
	router.Static("/js", STATIC_JS)
	router.Static("/css", STATIC_CSS)
	router.Static("/json", STATIC_JSON)
	router.Static("/tags", STATIC_TAGS)

	influxConfig := &configuration.InfluxConfig{}
	influxConfig.Init()
	influxConfig.Verify()

	// add route for index page
	homepageController := &controllers.HomepageController{}
	homepageController.Run(router)

	// add routes for bubbles
	bubbleController := &controllers.BubbleController{*influxConfig}
	bubbleController.Run(router)

	router.Run(":8080")
}
