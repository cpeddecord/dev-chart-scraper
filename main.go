package main

import (
	"fmt"
)

var rootURL = "https://www.digitaltruth.com/chart/"
var chartNotes = rootURL + "notes.php"
var printRoot = rootURL + "print.php"

func main() {
	indexHTML := GetHTML(printRoot);
	urls := ScrapeIndex(indexHTML)

	var h []HashMap
	for _, url := range urls {
		chartHTML := GetHTML(url)
		d, _ := ScrapeChart(chartHTML)

		h = append(h, d...)
	}

	fmt.Println(h)
}
