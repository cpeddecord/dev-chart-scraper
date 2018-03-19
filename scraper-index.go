package main

import (
	"io"
	"golang.org/x/net/html"
	"strings"
)

type strSlice []string

func ScrapeIndex(body io.ReadCloser) strSlice {
	defer body.Close()
	var s []string

	z := html.NewTokenizer(body)

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