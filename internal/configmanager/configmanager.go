package configmanager

import (
	"encoding/json"
	"io/ioutil"
)

type MysqlConnection struct {
	Host   string `json:"host"`
	DBName string `json:"dbname"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
}

type ApplicatonConfig struct {
	MySQL       MysqlConnection `json:"mysql"`
	ProcessName string          `json:"process_name"`
}

// Config values
var Config *ApplicatonConfig
var configFile *string

// GetConfig gtes the config values
func GetConfig() (ApplicatonConfig, error) {
	if Config != nil {
		return *Config, nil
	}
	err := LoadConfiguration()
	if err != nil {
		return ApplicatonConfig{}, err
	}
	return *Config, nil
}

// LoadConfiguration will initialize app config with config file name
func LoadConfiguration() error {
	var err error
	config := new(ApplicatonConfig)
	configFilepath := "config.json"
	raw, err := ioutil.ReadFile(configFilepath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw, config)
	if err != nil {
		return err
	}
	Config = config
	return nil
}
