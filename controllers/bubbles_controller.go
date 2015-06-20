package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"html/template"
	"github.com/influxdb/influxdb-viz/configuration"
)

const (
	TEMPLATE_BUBBLES = "views/bubbles.tmpl"
)

type BubbleController struct {
	InfluxConfig configuration.InfluxConfig
}

/* This method inits all the routes and deletegates to the associated resources component the gathering of data */
func (b *BubbleController) Run(router *gin.Engine) {

	bubblesResource := &BubblesResource{InfluxConfig:b.InfluxConfig}

	// main bubbles page
	router.GET("/bubbles", func(c *gin.Context) {
		obj := gin.H{"title": "HomePage"}
		router.SetHTMLTemplate(template.Must(template.ParseFiles(TEMPLATE_MAIN, TEMPLATE_BUBBLES)))
		c.HTML(http.StatusOK, "base", obj)
	})

	// return json flares format for d3js
	router.GET("/bubbles.json", bubblesResource.ListAll)
}

