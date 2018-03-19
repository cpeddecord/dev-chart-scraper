package main

import (
		"os"
		"testing"
)

// var _ io.Reader = (*os.File)(nil)
var chart, _ = os.Open("test/chart")
var nullChart, _ = os.Open("test/null-chart")

func TestScrapeChart(t *testing.T) {
	d, _ := ScrapeChart(chart)
	switch {
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