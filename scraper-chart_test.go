package main

import (
		"os"
		"testing"
		"golang.org/x/net/html"
)

var chartFile, _ = os.Open("test/chart")
var nullChartFile, _ = os.Open("test/null-chart")

var chart, _ = html.Parse(chartFile)
var nullChart, _ = html.Parse(nullChartFile)

func TestScrapeChart(t *testing.T) {

	d, err := ScrapeChart(chart)

	switch {
	case !err:
		t.Errorf("scraper returned truthy for some weird reason...")
	case d[0]["Film"] != "Adox CHM 125":
		t.Errorf("key Film didn't match Adox CHM 125")
	case d[1]["Developer"] != "510-Pyro":
		t.Errorf("key Developer didn't match")
	case d[2]["120"] != "":
		t.Errorf("key 120 should be blank")
	}

	_, ok := ScrapeChart(nullChart)
	if ok {
		t.Errorf("scraper returned data for some weird reason")
	}
}