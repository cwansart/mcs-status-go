package main

import (
	"fmt"

	"de.cwansart.mcss/settings"
	"de.cwansart.mcss/status"
)

func main() {
	url := settings.Get(settings.ServerUrlKey)
	status := status.Get(url)
	fmt.Printf("Player Count: %v\n", status.Players.Count)
}
