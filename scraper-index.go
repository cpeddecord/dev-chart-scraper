package main

import (
	"net/http"
	"golang.org/x/net/html"
	"strings"
)

type strSlice []string

func ScrapeIndex(url string) strSlice {
	var s []string

	res, _ := http.Get(printRoot)
	z := html.NewTokenizer(res.Body)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			return s

		case tt == html.StartTagToken:
			t:= z.Token()
			isAnchor := t.Data == "a"

			if isAnchor {
				for _, a := range t.Attr {
					if a.Key == "href" && strings.Contains(a.Val, "Film") {
						s = append(s, a.Val)
					}
				}
			}
		}
	}

	return s
}