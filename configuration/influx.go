package configuration

import (
	"net/url"
	"log"
	"fmt"
	"github.com/influxdb/influxdb/client"
	//"github.com/spf13/viper"
)

type InfluxConfig struct {
	InfluxDbHost string
	InfluxDbPort int
	InfluxDbUser string
	InfluxDbPassword string
}

func (Config *InfluxConfig) Init() {

	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/init.d/influxdb-viz/")   // path to look for the config file in
	viper.AddConfigPath("$HOME/.influxdb-viz")  // call multiple times to add many search paths
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	Config.InfluxDbHost = "localhost"
	Config.InfluxDbPort = 8086
	Config.InfluxDbPassword = "root"
	Config.InfluxDbUser = "root"

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
		log.Fatalf("Impossible to ping InfluxDB : %s", err)
	}
	log.Printf("Happy to be connected to InfluxDB! %v, %s", dur, ver)
}


