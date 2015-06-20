// The MIT License (MIT)
//
// Copyright (c) 2015 Nicolas MULLER (@zepouet)
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package configuration

import (
	"net/url"
	"log"
	"fmt"
	"github.com/influxdb/influxdb/client"
	"github.com/spf13/viper"
	"runtime"
)

type InfluxConfig struct {
	InfluxDbHost string
	InfluxDbPort int
	InfluxDbUser string
	InfluxDbPassword string
}

func (Config *InfluxConfig) Init() {

	// The configuration file will be named config.toml
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	// We could search the configuration file into these directories
	viper.AddConfigPath("$HOME/.influxdb-viz")
	viper.AddConfigPath("$HOME/settings")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		fmt.Println("No configuration file loaded - using defaults")
		fmt.Println("Fatal error config file: %s \n", err)
	}

	log.Println("InfluxDbHost " + viper.GetString("InfluxDbHost"))
	log.Println("InfluxDbPort " + viper.GetString("InfluxDbPort"))
	log.Println("InfluxDbUser " + viper.GetString("InfluxDbUser"))
	log.Println("InfluxDbPassword " + viper.GetString("InfluxDbPassword"))

	Config.InfluxDbHost = viper.GetString("InfluxDbHost")
	Config.InfluxDbPort = viper.GetInt("InfluxDbPort")
	Config.InfluxDbUser = viper.GetString("InfluxDbUser")
	Config.InfluxDbPassword = viper.GetString("InfluxDbPassword")

	// Halt and catch fire
	runtime.GOMAXPROCS(runtime.NumCPU())
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



