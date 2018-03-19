package main

import (
	"golang.org/x/net/html"
	"bytes"
	"io"
)

var RootURL = "https://www.digitaltruth.com/chart/"
var ChartNotes = RootURL + "notes.php"
var PrintRoot = RootURL + "print.php"

func RenderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}