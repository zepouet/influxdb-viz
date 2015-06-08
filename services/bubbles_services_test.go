package services

import (
	"fmt"
	"time"
	"testing"
	"net/url"
	"math/rand"
	"github.com/influxdb/influxdb/client"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {

	// setup (run before each `Convey` at this scope):
	fmt.Println("Connect to the InfluxDB Server")

	Convey("Play with the bubbles", t, func() {

		// setup (run before each `Convey` at this scope):
		fmt.Println("Create a new database")
		host, err := url.Parse(fmt.Sprintf("http://%s:%d", "localhost", 8086))
		So(err, ShouldBeNil)
		con, err := client.NewClient(client.Config{URL: *host})
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
					Database:        "BumbeBeeTuna",
					RetentionPolicy: "default",
				}
				_, err = con.Write(bps)
				So(err, ShouldBeNil)

				Convey("Select the series", func() {
					con, err := client.NewClient(client.Config{URL: *host})
					q := client.Query{
						Command:  "select count(value) from shapes",
						Database: "BumbeBeeTuna",
					}
					response, err := con.Query(q);
					if err == nil && response.Error() == nil {
						//fmt.Println(response.Results.UnmarshalJSON())
					}
				})

			})

		})

		Reset(func() {
			q := client.Query{
				Command:  "delete from shapes",
				Database: "BumbeBeeTuna",
			}
			response, err := con.Query(q);
			if err == nil && response.Error() == nil {
				fmt.Println(response.Results)
			}
		})

	})

}
