package main

import (
	"fmt"
	"io"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type strArr []string
type HashMap map[string]string
type mapSlice []HashMap

func thMatcher(n *html.Node) bool {
	if n.DataAtom == atom.Th && n.Parent != nil && n.Parent.Parent != nil {
		return scrape.Text(n) != ""
	}

	return false
}

func trMatcher(n *html.Node) bool {
	if n.DataAtom == atom.Tr && n.Parent != nil && n.Parent.Parent != nil {
		return scrape.Text(n) != ""
	}

	return false
}

func getRowHTMLs(n *html.Node) strArr {
	var ret strArr
	var current = n

	for {
		if current.NextSibling == nil {
			break
		}

		v := ""
		if current.FirstChild != nil && current.FirstChild.Data != "" {
			v = current.FirstChild.Data
		}
		ret = append(ret, v)
		current = current.NextSibling
	}

	return ret
}

func ScrapeChart(body io.ReadCloser) (mapSlice, bool) {
	defer body.Close()
	var s []HashMap

	root, err := html.Parse(body)
	if err != nil {
		panic(err)
	}


	// there's no table head, bail
	thNode, ok := scrape.Find(root, thMatcher)
	if (ok != true) {
		return s, false
	}
	keys := getRowHTMLs(thNode)

	var d []strArr
	trNodes := scrape.FindAll(root, trMatcher)
	for i, n := range trNodes {
		if i == 0 { continue }
		d = append(d, getRowHTMLs(n.FirstChild))
	}

	for _, vals := range d {
		m := make(HashMap)
		for i, v := range vals {
			m[keys[i]] = v
		}
		s = append(s, m)
	}

	fmt.Println(s)
	return s, true
}
