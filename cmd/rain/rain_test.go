package main

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"testing"
	"time"
)

func TestDisplay(t *testing.T) {

	body, err := ioutil.ReadFile("../../test/fixture.xml")
	if err != nil {
		t.Errorf("Could not read file")
	}
	w := Weatherdata{}
	err = xml.Unmarshal(body, &w)
	if err != nil {
		t.Errorf("Could not marshall xml")
	}
	buf := new(bytes.Buffer)
	location, err := time.LoadLocation("Europe/Dublin")
	if err != nil {
		panic(err)
	}
	now := time.Date(
		2019, 5, 9, 12, 34, 58, 651387237, location)
	display(buf, w, 12, now, "Europe/Dublin")
	expected := `12:00 - 13:00 0.0 mm | 
13:00 - 14:00 0.0 mm | 
14:00 - 15:00 0.0 mm | 
15:00 - 16:00 0.0 mm | 
16:00 - 17:00 0.1 mm | *
17:00 - 18:00 0.0 mm | 
18:00 - 19:00 0.0 mm | 
19:00 - 20:00 0.0 mm | 
20:00 - 21:00 0.0 mm | 
21:00 - 22:00 0.0 mm | 
22:00 - 23:00 0.0 mm | 
23:00 - 00:00 0.0 mm | 
`
	if buf.String() != expected {
		t.Errorf("Output mismatch. Expected %s, Actual %s",
			expected, buf.String())
	}
}

func TestDisplaySkipFirstHour(t *testing.T) {

	body, err := ioutil.ReadFile("../../test/fixture.xml")
	if err != nil {
		t.Errorf("Could not read file")
	}
	w := Weatherdata{}
	err = xml.Unmarshal(body, &w)
	if err != nil {
		t.Errorf("Could not marshall xml")
	}
	buf := new(bytes.Buffer)
	location, err := time.LoadLocation("Europe/Dublin")
	if err != nil {
		panic(err)
	}
	now := time.Date(
		2019, 5, 9, 13, 34, 58, 651387237, location)
	display(buf, w, 18, now, "Europe/Dublin")
	expected := `13:00 - 14:00 0.0 mm | 
14:00 - 15:00 0.0 mm | 
15:00 - 16:00 0.0 mm | 
16:00 - 17:00 0.1 mm | *
17:00 - 18:00 0.0 mm | 
18:00 - 19:00 0.0 mm | 
19:00 - 20:00 0.0 mm | 
20:00 - 21:00 0.0 mm | 
21:00 - 22:00 0.0 mm | 
22:00 - 23:00 0.0 mm | 
23:00 - 00:00 0.0 mm | 
00:00 - 01:00 0.0 mm | 
01:00 - 02:00 0.0 mm | 
02:00 - 03:00 0.0 mm | 
03:00 - 04:00 0.0 mm | 
04:00 - 05:00 0.0 mm | 
05:00 - 06:00 0.0 mm | 
06:00 - 07:00 0.0 mm | 
`
	if buf.String() != expected {
		t.Errorf("Output mismatch. Expected %s, Actual %s",
			expected, buf.String())
	}
}

func TestDisplayNonHourTimezone(t *testing.T) {

	body, err := ioutil.ReadFile("../../test/fixture.xml")
	if err != nil {
		t.Errorf("Could not read file")
	}
	w := Weatherdata{}
	err = xml.Unmarshal(body, &w)
	if err != nil {
		t.Errorf("Could not marshall xml")
	}
	buf := new(bytes.Buffer)
	location, err := time.LoadLocation("Asia/Rangoon")
	if err != nil {
		panic(err)
	}
	now := time.Date(
		2019, 5, 9, 13, 34, 58, 651387237, location)
	display(buf, w, 18, now, "Asia/Rangoon")
	expected := `17:30 - 18:30 0.0 mm | 
18:30 - 19:30 0.0 mm | 
19:30 - 20:30 0.0 mm | 
20:30 - 21:30 0.0 mm | 
21:30 - 22:30 0.1 mm | *
22:30 - 23:30 0.0 mm | 
23:30 - 00:30 0.0 mm | 
00:30 - 01:30 0.0 mm | 
01:30 - 02:30 0.0 mm | 
02:30 - 03:30 0.0 mm | 
03:30 - 04:30 0.0 mm | 
04:30 - 05:30 0.0 mm | 
05:30 - 06:30 0.0 mm | 
06:30 - 07:30 0.0 mm | 
07:30 - 08:30 0.0 mm | 
08:30 - 09:30 0.0 mm | 
09:30 - 10:30 0.0 mm | 
10:30 - 11:30 0.0 mm | 
`
	if buf.String() != expected {
		t.Errorf("Output mismatch. Expected %s, Actual %s",
			expected, buf.String())
	}
}
