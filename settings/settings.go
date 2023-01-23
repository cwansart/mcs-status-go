package settings

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

const (
	configFile = "config.json"
	configPath = "."
)

type Config struct {
	ServerUrl string `json:"serverurl"`
}

var config = Config{
	ServerUrl: "http://localhost:2006",
}

func ReadConfigFile() Config {
	log.Println("Loading config file")
	p := fmt.Sprintf("%v/%v", configPath, configFile)

	if _, err := os.Stat(p); errors.Is(err, os.ErrNotExist) {
		log.Printf("Could not find config file, generating one using defaults in %v", p)
		b, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatal("Could not marshal config data to json")
		}
		os.WriteFile(p, b, 0600)
	} else {
		b, err := os.ReadFile(p)
		if err != nil {
			log.Fatalf("Could not open config file %v", p)
		}
		json.Unmarshal(b, &config)
	}
	log.Println("Successfully read config file", p, config.ServerUrl)

	return config
}
