package status

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type Status struct {
	Players Players `json:"players"`
}

type Players struct {
	Count int32 `json:"count"`
}

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

func getResponseBody(r *http.Response) []byte {
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return b
}

func Get(url string) (s Status) {
	r := getServerStatus(url)
	b := getResponseBody(r)
	json.Unmarshal(b, &s)
	return
}
