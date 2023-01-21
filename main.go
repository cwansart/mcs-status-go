package main

import (
	"fmt"

	"de.cwansart.mcss/settings"
	"de.cwansart.mcss/status"
)

func main() {
	url := settings.Get(settings.ServerUrlKey)
	status := status.Get(url)
	fmt.Printf("status: %v\n", status)
}
