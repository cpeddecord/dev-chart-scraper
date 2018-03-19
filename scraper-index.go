package main

import (
	"strings"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type strSlice []string

func anchorMatcher(n *html.Node) bool {
	if n.DataAtom == atom.A {
		a := scrape.Attr(n, "href")
		return strings.Contains(a, "Film")
	}

	return false
}

func ScrapeIndex(h *html.Node) (strSlice) {
	var s strSlice

	aNodes := scrape.FindAll(h, anchorMatcher)

	for _, anchor := range aNodes {
		for _, attr := range anchor.Attr {
			s = append(s, attr.Val)
		}
	}

	return s
}