package main

import (
	"fmt"
)

var rootURL = "https://www.digitaltruth.com/chart/"
var chartNotes = rootURL + "notes.php"
var printRoot = rootURL + "print.php"

func main() {
	urls := ScrapeIndex(printRoot)

	fmt.Println(urls)
}