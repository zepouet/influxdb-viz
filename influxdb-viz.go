// The MIT License (MIT)
//
// Copyright (c) 2015 Nicolas MULLER (@zepouet)
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

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

	// read the configuration and exits if necessary
	influxConfig := &configuration.InfluxConfig{}
	influxConfig.Init()
	influxConfig.Verify()

	// initialise the web engine
	router := gin.Default()

	// create routes for all static files
	router.Static("/images", STATIC_IMAGES)
	router.Static("/js", STATIC_JS)
	router.Static("/css", STATIC_CSS)
	router.Static("/json", STATIC_JSON)
	router.Static("/tags", STATIC_TAGS)

	// add route for index page
	homepageController := &controllers.HomepageController{}
	homepageController.Run(router)

	// add routes for bubbles
	bubbleController := &controllers.BubbleController{*influxConfig}
	bubbleController.Run(router)

	router.Run(":8080")
}
