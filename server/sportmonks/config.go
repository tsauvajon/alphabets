package sportmonks

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

type config struct {
	Token string `json:"token"`
}

func getConfig() (config, error) {
	var cfg config

	pwd, err := os.Getwd()

	if err != nil {
		fmt.Printf("Error getting working dir: %v\n", err)
		os.Exit(1)
	}

	secretPath := path.Join(pwd, "sportmonks", "secret.json")

	fmt.Println(secretPath)

	file, err := os.Open(secretPath)

	if err != nil {
		return config{}, fmt.Errorf("Error opening secret.json: %v", err)
	}

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&cfg)

	if err != nil {
		return config{}, fmt.Errorf("Error decoding the config file: %v", err)
	}

	return cfg, nil
}
