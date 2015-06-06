package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/influxdb/influxdb-viz/services"
	"github.com/influxdb/influxdb-viz/configuration"
)

type BubblesResource struct {
	InfluxConfig configuration.InfluxConfig
}

func (br *BubblesResource) ListAll(c *gin.Context) {
	service := &services.BubblesService{InfluxConfig:br.InfluxConfig}
	bubbles := service.ListAll()
	c.JSON(200, bubbles)
}
