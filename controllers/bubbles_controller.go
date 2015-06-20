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

