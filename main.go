package main

import (
	"encoding/json"
	"fmt"

	"de.cwansart.mcss/settings"
	"de.cwansart.mcss/status"
)

func main() {
	url := settings.Get(settings.ServerUrlKey)
	r := status.Get(url)

	json, _ := json.Marshal(r)

	fmt.Printf("Player Count: %v\n", string(json))
}
