package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"de.cwansart.mcss/settings"
)

func getHttpClient() http.Client {
	return http.Client{
		Timeout: 3 * time.Second,
	}
}

func getServerStatus(url string) *http.Response {
	log.Printf("Calling %v", url)
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
	settings.ReadConfig()
	mcUrl := settings.Get(settings.ServerUrlKey)
	readAndPrintData(mcUrl)
}
