package configger

import (
	"encoding/json"
	"os"
	"visiontest/infrastructure/logger"
)

type Config struct {
	Appname  string `json:"appname"`
	Hostport string `json:"hostport"`
}

var Conf *Config

func init() {
	Conf, _ = LoadConfig()
}

func LoadConfig() (*Config, error) {
	data, err := os.ReadFile("conf/appconf.json")
	if err != nil {
		logger.Error("configger LoadConfig  Error reading config file")
		return nil, err
	}
	Conf := new(Config)
	err = json.Unmarshal(data, Conf)
	if err != nil {
		return nil, err
	}
	return Conf, nil
}
