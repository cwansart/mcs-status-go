package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func getHttpClient() http.Client {
	return http.Client{
		Timeout: 3 * time.Second,
	}
}

func getServerStatus(url string) *http.Response {
	c := getHttpClient()
	r, err := c.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	return r
}

func getResponseBody(r *http.Response) string {
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
		return ""
	}
	return string(b)
}

func readAndPrintData(url string) {
	r := getServerStatus(url)
	b := getResponseBody(r)

	fmt.Println(b)
}

func main() {
	mcUrl := "http://localhost:8080/"
	readAndPrintData(mcUrl)
}
