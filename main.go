package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"bytes"
	"io"
)

var rootURL = "https://www.digitaltruth.com/chart/"
var chartNotes = rootURL + "notes.php"
var printRoot = rootURL + "print.php"

func getHTML(url string) (*html.Node, error) {
	fmt.Println("GET ", url)
	res, err := http.Get(printRoot)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	return html.Parse(res.Body)
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}

func main() {
	indexHTML, err := getHTML(printRoot);
	if (err != nil) {
		fmt.Println("welp, time to go home")
		return
	}
	// defer printRoot
	urls:= ScrapeIndex(indexHTML)

	var h []HashMap
	for _, u := range urls {
		chartHTML, _ := getHTML(rootURL + u)
		htmls := renderNode(chartHTML)
		fmt.Println(htmls)
		d, ok := ScrapeChart(chartHTML)
		if ok {
			h = append(h, d...)
		}
	}

	fmt.Println(h)
}
