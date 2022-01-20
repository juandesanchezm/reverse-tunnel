package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Load loads a YAML document from a file and fills in corresponding fields of
// conf variable. Default values given in the variable are honored.
func Load(path string, conf interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, conf)
}

// Load loads a YAML document from a string and fills in corresponding fields of
// conf variable. Default values given in the variable are honored.
func LoadString(data string, conf interface{}) error {
	return yaml.Unmarshal([]byte(data), conf)
}
