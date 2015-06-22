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

package services

import (
	"fmt"
	"log"
	"net/url"
	"github.com/influxdb/influxdb-viz/models"
	"github.com/influxdb/influxdb-viz/configuration"
	"github.com/influxdb/influxdb/client"
)

type SeriesServices struct {
	InfluxConfig configuration.InfluxConfig
}

func (br *SeriesServices) ListAll() models.Serie {
	serie := models.Serie{Name:"shapes"}

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
		Database: "testing",
	}
	response, err := con.Query(q)
	if err == nil && response.Error() == nil {
		res := response.Results
		log.Println(res[0])
		//count := res[0].Series[0].Values[0][1]
	} else {
		log.Println(err)
	}

	return serie
}
