package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	//"github.com/spf13/viper"
)

const (
	TEMPLATES        = "templates/"
	TEMPLATE_MAIN    = TEMPLATES + "main.tmpl"
	TEMPLATE_INDEX   = TEMPLATES + "index.tmpl"
	TEMPLATE_BUBBLES = TEMPLATES + "bubbles.tmpl"
	STATIC           = "static/"
	STATIC_JS        = STATIC + "js"
	STATIC_CSS       = STATIC + "css"
	STATIC_IMAGES    = STATIC + "images"
)

func main() {
	router := gin.Default()

	// create static route for all files in this folder
	router.Static("/images", STATIC_IMAGES)
	router.Static("/js", STATIC_JS)
	router.Static("/css", STATIC_CSS)

	// index page
	router.GET("/", func(c *gin.Context) {
		obj := gin.H{"title": "Index"}
		router.SetHTMLTemplate(template.Must(template.ParseFiles(TEMPLATE_MAIN, TEMPLATE_INDEX)))
		c.HTML(200, "base", obj)
	})

	// bubbles page
	router.GET("/bubbles", func(c *gin.Context) {
		obj := gin.H{"title": "Index"}
		router.SetHTMLTemplate(template.Must(template.ParseFiles(TEMPLATE_MAIN, TEMPLATE_BUBBLES)))
		c.HTML(200, "base", obj)
	})

	router.Run(":8080")
}
