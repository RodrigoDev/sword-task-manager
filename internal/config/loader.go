package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

const fileName = "config.yaml"

func GetConfig() (*ServiceConfig, error) {
	var config ServiceConfig

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}