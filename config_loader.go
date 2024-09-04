package configloader

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

func LoadJSONConfig(filePath string, config interface{}) error {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, config)
	if err != nil {
		return err
	}
	return nil
}

func LoadYAMLConfig(filePath string, config interface{}) error {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		return err
	}
	return nil
}

func LoadENVConfig(filePath string) error {
	err := godotenv.Load(filePath)
	if err != nil {
		return errors.New("Error loading .env file")
	}
	return nil
}
