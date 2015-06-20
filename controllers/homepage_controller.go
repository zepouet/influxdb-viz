package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"html/template"
)

const (
	TEMPLATE_INDEX   = "templates/index.tmpl"
	TEMPLATE_MAIN    = "templates/main.tmpl"
)

type HomepageController struct {
}

func (b *HomepageController) Run(router *gin.Engine) {
	// bubbles page
	router.GET("/", func(c *gin.Context) {
		obj := gin.H{"title": "HomePage"}
		router.SetHTMLTemplate(template.Must(template.ParseFiles(TEMPLATE_MAIN, TEMPLATE_INDEX)))
		c.HTML(http.StatusOK, "base", obj)
	})
}
