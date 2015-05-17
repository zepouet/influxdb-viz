package bubbles

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"html/template"
)

const (
	TEMPLATE_MAIN    =  "templates/main.tmpl"
	TEMPLATE_BUBBLES = "templates/bubbles.tmpl"
)

type Config struct {
	InfluxDbHost string
	InfluxDbPort string
	InfluxDbUser string
	InfluxDbPassword string
}

type BubbleService struct {
}

func (b *BubbleService) Run(config Config) {
	router := gin.Default()

	// bubbles page
	router.GET("/bubbles", func(c *gin.Context) {
		obj := gin.H{"title": "Bubbles"}
		router.SetHTMLTemplate(template.Must(template.ParseFiles(TEMPLATE_MAIN, TEMPLATE_BUBBLES)))
		c.HTML(http.StatusOK, "base", obj)
	})
}