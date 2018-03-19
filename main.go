package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func getHTML(url string) (*html.Node, error) {
	fmt.Println("GET ", url)
	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	return html.Parse(res.Body)
}

func main() {
	indexHTML, err := getHTML(PrintRoot);
	if (err != nil) {
		fmt.Println("welp, time to go home")
		return
	}
	urls:= ScrapeIndex(indexHTML)

	var h []HashMap
	for _, u := range urls {
		chartHTML, _ := getHTML(RootURL + u)
		d, ok := ScrapeChart(chartHTML)
		if ok {
			h = append(h, d...)
		}
	}

	fmt.Println(h)
}
