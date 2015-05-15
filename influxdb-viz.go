package main

import (
	"net/http"
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
	STATIC_JSON    = STATIC + "json"
)

func main() {
	router := gin.Default()

	// create routes for all static files
	router.Static("/images", STATIC_IMAGES)
	router.Static("/js", STATIC_JS)
	router.Static("/css", STATIC_CSS)
	router.Static("/json", STATIC_JSON)

	// index page
	router.GET("/", func(c *gin.Context) {
		obj := gin.H{"title": "Index"}
		router.SetHTMLTemplate(template.Must(template.ParseFiles(TEMPLATE_MAIN, TEMPLATE_INDEX)))
		c.HTML(http.StatusOK, "base", obj)
	})

	// bubbles page
	router.GET("/bubbles", func(c *gin.Context) {
		obj := gin.H{"title": "Bubbles"}
		router.SetHTMLTemplate(template.Must(template.ParseFiles(TEMPLATE_MAIN, TEMPLATE_BUBBLES)))
		c.HTML(http.StatusOK, "base", obj)
	})

	router.Run(":8080")
}
