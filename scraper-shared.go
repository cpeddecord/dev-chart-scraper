package main

import (
	"fmt"
	"io"
	"net/http"
)

func GetHTML(url string) io.ReadCloser {
	res, err := http.Get(printRoot)

	if err != nil {
		fmt.Println(err)
	}

	return res.Body
}