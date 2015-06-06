package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"html/template"
)

const (
	TEMPLATE_MAIN    = TEMPLATES + "main.tmpl"
	TEMPLATE_BUBBLES = TEMPLATES + "bubbles.tmpl"
)

type Config struct {
	InfluxDbHost string
	InfluxDbPort string
	InfluxDbUser string
	InfluxDbPassword string
}

type BubbleController struct {
}

func (b *BubbleController) Run(router *gin.Engine) {
	// bubbles page
	router.GET("/bubbles", func(c *gin.Context) {
		obj := gin.H{"title": "Bubbles"}
		router.SetHTMLTemplate(template.Must(template.ParseFiles(TEMPLATE_MAIN, TEMPLATE_BUBBLES)))
		c.HTML(http.StatusOK, "base", obj)
	})
}