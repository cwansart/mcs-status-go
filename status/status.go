package status

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type status struct {
	Players players `json:"players"`
}

type players struct {
	Count int32 `json:"count"`
}

type Response struct {
	IsOnline    bool  `json:"online"`
	PlayerCount int32 `json:"player_count"`
}

func getHttpClient() http.Client {
	return http.Client{
		Timeout: 500 * time.Millisecond,
	}
}

func getServerStatus(url string) (*http.Response, error) {
	log.Printf("Calling %v", url)
	c := getHttpClient()
	r, err := c.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return r, nil
}

func getResponseBody(r *http.Response) ([]byte, error) {
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return b, nil
}

func getStatus(url string) (s status, err error) {
	r, err := getServerStatus(url)
	if err != nil {
		return status{}, err
	}

	b, err := getResponseBody(r)
	if err != nil {
		return status{}, err
	}

	err = json.Unmarshal(b, &s)
	if err != nil {
		return status{}, err
	}

	return s, nil
}

func Get(url string) Response {
	r := Response{}
	s, err := getStatus(url)

	if err == nil {
		r.PlayerCount = s.Players.Count
		r.IsOnline = true
	}

	return r
}
