package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type LocationType struct {
	Precipitation Precipitation `xml:" precipitation,omitempty"`
}

type Precipitation struct {
	Unit  string  `xml:"unit,attr"`
	Value float64 `xml:"value,attr"`
}

type ProductType struct {
	Time []TimeType `xml:" time"`
}

type TimeType struct {
	Location []LocationType `xml:" location"`
	From     time.Time      `xml:"from,attr"`
	To       time.Time      `xml:"to,attr"`
}

type Unitvalue struct {
	Unit  string  `xml:"unit,attr"`
	Value float32 `xml:"value,attr"`
}

type Weatherdata struct {
	Product []ProductType `xml:" product,omitempty"`
}

func display(f io.Writer, w Weatherdata, hours int,
	now time.Time, zone string) {
	current := 0
	tz, err := time.LoadLocation(zone)
	if err != nil {
		log.Fatal(`Failed to load location "Local" for timezone`)
	}
	for _, t := range w.Product[0].Time {
		// Filter out longer time difference data
		if t.To.Sub(t.From).Hours() != 1 {
			continue
		}
		// Filter out anything stale (predicting the past)
		if t.To.In(tz).Sub(now).Hours() < 0 {
			continue
		}
		for _, l := range t.Location {
			stars := strings.Repeat("*", int(l.Precipitation.Value*10.0))
			fmt.Fprintf(f, "%02d:%02d - %02d:%02d %.1f mm | %s\n",
				t.From.In(tz).Hour(), t.From.In(tz).Minute(),
				t.To.In(tz).Hour(), t.To.In(tz).Minute(),
				l.Precipitation.Value,
				stars,
			)
			current++
			if current == hours {
				return
			}
		}
	}
}

func main() {

	w := Weatherdata{}

	hours := flag.Int("hours", 12, "Number of hours to forecast")
	latitude := flag.Float64(
		"latitude", 53.292148,
		"Latitude. Use to specify forecast location",
	)
	longitude := flag.Float64(
		"longitude", -9.007064,
		"Longitude. Use to specify forecast location",
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s: Rain forecast\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  Note: Single or double dash can be used "+
			"for parameters, eg -hours/--hours.\n\n")
		flag.PrintDefaults()
	}

	flag.Parse()
	resp, err := http.Get(
		"https://api.met.no/weatherapi/locationforecast/1.9/" +
			"?lat=" + fmt.Sprintf("%f", *latitude) +
			";lon=" + fmt.Sprintf("%f", *longitude),
	)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = xml.Unmarshal(body, &w)
	if err != nil {
		log.Fatal(err)
	}

	t := time.Now()
	display(os.Stdout, w, *hours, t, "Local")
}
