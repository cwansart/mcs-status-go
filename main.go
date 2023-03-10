package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cwansart/mcs-status-go/config"
	"github.com/cwansart/mcs-status-go/status"
)

func main() {
	c := config.NewConfig("./config.json")
	url := c.ServerUrl

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s := status.Get(url)
		b, err := json.Marshal(s)

		if err != nil {
			log.Println("Failed to convert status to json:", s)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 internal server error"))
		} else {
			log.Println("Received request, response:", string(b))
			w.Header().Add("Content-Type", "application/json")
			w.Write(b)
		}
	})

	// http.ListenAndServeTLS()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
