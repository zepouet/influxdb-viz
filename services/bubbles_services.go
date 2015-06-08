package services

import (
	"fmt"
	"log"
	"net/url"
	"github.com/influxdb/influxdb-viz/models"
	"github.com/influxdb/influxdb-viz/configuration"
	"github.com/influxdb/influxdb/client"
)

type BubblesService struct {
	InfluxConfig configuration.InfluxConfig
}

func (br *BubblesService) ListAll() models.Bubble {
	bubbleFlare := models.Bubble{Name:"flare"}

	fmt.Printf("http://%s:%d", br.InfluxConfig.InfluxDbHost, br.InfluxConfig.InfluxDbPort)

	host, err := url.Parse(fmt.Sprintf("http://%s:%d", br.InfluxConfig.InfluxDbHost, br.InfluxConfig.InfluxDbPort))
	if err != nil {
		log.Fatal(err)
	}
	con, err := client.NewClient(client.Config{URL: *host, Username:"root", Password:"root"})
	if err != nil {
		log.Fatal(err)
	}

	q := client.Query{
		Command:  "show series",
		Database: "demo",
	}
	response, err := con.Query(q)
	if err == nil && response.Error() == nil {
		log.Println(response.Results)
	} else {
		log.Println(err)
	}

	bubble1 := models.Bubble{Name:"serie1", Size:13,}
	bubble2 := models.Bubble{Name:"serie2", Size:3,}
	bubble3 := models.Bubble{Name:"serie3", Size:17,}
	bubbleFlare.Children = append(bubbleFlare.Children, bubble1, bubble2, bubble3)
	return bubbleFlare
}
