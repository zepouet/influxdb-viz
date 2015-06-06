package bubbles

import (
	"github.com/influxdb/influxdb/client"
	"github.com/gin-gonic/gin"
)

type BubbleResource struct {
	clientInflux client.Client
}

func (br *BubbleResource) listSeries(c *gin.Context) {
}
