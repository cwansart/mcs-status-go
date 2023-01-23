package settings

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

const configFile = "./config.json"

type Config struct {
	ServerUrl string `json:"serverurl"`
}

var config = Config{
	ServerUrl: "http://localhost:2006",
}

func createConfigFile(p string) {
	log.Printf("Could not find config file, generating one using defaults in %v", p)
	b, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatal("Could not marshal config data to json")
	}
	os.WriteFile(p, b, 0600)
}

func loadConfigFile(p string) {
	b, err := os.ReadFile(p)
	if err != nil {
		log.Fatalf("Could not open config file %v", p)
	}
	json.Unmarshal(b, &config)
}

func ReadConfigFile() Config {
	log.Println("Loading config file")

	if _, err := os.Stat(configFile); errors.Is(err, os.ErrNotExist) {
		createConfigFile(configFile)
	} else {
		loadConfigFile(configFile)
	}
	log.Println("Successfully read config file", configFile, config.ServerUrl)

	return config
}
