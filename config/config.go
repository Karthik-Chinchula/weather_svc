package config

import (
	"encoding/json"
	"os"
	"weather_svc/models"
)

// ReadConfig will read the config file from given path and load server configs.
func ReadConfig(filePath string) (models.Config, error) {
	var config models.Config

	// Open the config file
	configFile, err := os.Open(filePath)
	if err != nil {
		return config, err
	}
	defer configFile.Close()

	// Decode the config file
	if err = json.NewDecoder(configFile).Decode(&config); err != nil {
		return config, err
	}

	return config, nil
}
