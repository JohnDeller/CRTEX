package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Database struct {
		Host     string `json:"host,omitempty"`
		Port     string `json:"port,omitempty"`
		UserName string `json:"user_name,omitempty"`
		DBName   string `json:"db_name,omitempty"`
		SSLMode  string `json:"ssl_mode,omitempty"`
	} `json:"database,omitempty"`

	Broker struct {
		Host  string `json:"host,omitempty"`
		Port  string `json:"port,omitempty"`
		User  string `json:"user,omitempty"`
		Queue struct {
			Name       string `json:"name"`
			Durable    bool   `json:"durable"`
			AutoDelete bool   `json:"auto_delete"`
			Exclusive  bool   `json:"exclusive"`
			NoWait     bool   `json:"no_wait"`
		} `json:"queue"`
	} `json:"broker,omitempty"`

	Host string `json:"host,omitempty"`
	Port string `json:"port,omitempty"`
}

func InitConfig(configFileName string) (Config, error) {
	var config Config
	configFile, err := os.Open(configFileName)
	defer configFile.Close()

	if err != nil {
		return config, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)

	return config, err
}
