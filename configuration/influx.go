package configuration

import (
	"net/url"
	"log"
	"fmt"
	"github.com/influxdb/influxdb/client"
)

type InfluxConfig struct {
	InfluxDbHost string
	InfluxDbPort int
	InfluxDbUser string
	InfluxDbPassword string
}

func (config *InfluxConfig) Verify() {
	host, err := url.Parse(fmt.Sprintf("http://%s:%d", config.InfluxDbHost, config.InfluxDbPort))
	if err != nil {
		log.Fatal(err)
	}
	con, err := client.NewClient(client.Config{URL: *host})
	if err != nil {
		log.Fatal(err)
	}

	dur, ver, err := con.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Happy to be connected to InfluxDB! %v, %s", dur, ver)
}

