package settings

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

const (
	configFile = "config.yaml"
	configName = "config"
	configType = "yaml"
	configPath = "."
)

const (
	ServerUrlKey = "serverurl"
)

func setDefaults() {
	viper.SetConfigFile(configFile)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)
	viper.SetDefault(ServerUrlKey, "http://localhost:2008")
}

func readConfigFile() {
	log.Println("Loading config file")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			newConfigFilePath := fmt.Sprintf("%v/%v", configPath, configFile)
			log.Printf("Could not find config file, generating one using defaults in %v", newConfigFilePath)
			viper.WriteConfigAs(newConfigFilePath)
		} else {
			log.Println("Could not load config file, using defaults.", err)
		}
	} else {
		log.Println("Successfully read config file", viper.ConfigFileUsed())
	}
}

func init() {
	setDefaults()
	readConfigFile()
}

func Get(key string) string {
	// TODO: return error on failure instead of calling log.Fatalf
	value := viper.GetString(key)
	if len(value) == 0 {
		log.Fatalf("Cannot find setting for key: '%v'", key)
	}
	return value
}
