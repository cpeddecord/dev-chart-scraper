package main

import (
		"os"
		"testing"
		"golang.org/x/net/html"
)

var indexFile, _ = os.Open("test/index")
var index, _ = html.Parse(indexFile)

func TestScrapeIndex(t *testing.T) {
	d := ScrapeIndex(index)
	switch {
	case d[0] != "search_text.php?Film=Adox+CHM":
		t.Error("test failed to get proper results")
	}
}
