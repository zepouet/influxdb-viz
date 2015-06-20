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
	"time"
	"testing"
	"net/url"
	"math/rand"
	"github.com/influxdb/influxdb/client"
	. "github.com/smartystreets/goconvey/convey"
)

// queryDB convenience function to query the database
func queryDB(con *client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: "testing",
	}
	if response, err := con.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	}
	return
}

func TestSpec(t *testing.T) {

	// setup (run before each `Convey` at this scope):
	fmt.Println("Connect to the InfluxDB Server")

	Convey("Play with the bubbles", t, func() {
		
		// setup (run before each `Convey` at this scope):
		fmt.Println("Create a new database")
		host, err := url.Parse(fmt.Sprintf("http://%s:%d", "localhost", 8086))
		con, err := client.NewClient(client.Config{URL: *host})
		_, err = queryDB(con, fmt.Sprintf("create database %s", "testing"))

		So(err, ShouldBeNil)

		Convey("Create many series", func() {

			var (
				shapes     = []string{"circle", "rectangle", "square", "triangle"}
				colors     = []string{"red", "blue", "green"}
				sampleSize = 10
				pts        = make([]client.Point, sampleSize)
			)

			Convey("Test inserts", func() {

				rand.Seed(42)
				for i := 0; i < sampleSize; i++ {
					pts[i] = client.Point{
						Measurement: "shapes",
						Tags: map[string]string{
							"color": colors[rand.Intn(len(colors))],
							"shape": shapes[rand.Intn(len(shapes))],
						},
						Fields: map[string]interface{}{
							"value": rand.Intn(sampleSize),
						},
						Time:      time.Now(),
						Precision: "s",
					}
				}

				bps := client.BatchPoints{
					Points:          pts,
					Database:        "testing",
					RetentionPolicy: "default",
				}
				_, err = con.Write(bps)
				So(err, ShouldBeNil)

				Convey("Select the series", func() {
					q := client.Query{
						Command:  "select count(value) from shapes",
						Database: "testing",
					}

					if response, err := con.Query(q); err == nil && response.Error() == nil {
						res := response.Results
						count := res[0].Series[0].Values[0][1]
						log.Println(count)
					}
				})

			})

		})

		Reset(func() {

		})

	})

}
