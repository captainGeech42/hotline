package config

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func LoadConfig(path string) *Config {
	config := &Config{}

	// check if $HOTLINE_CONFIG_PATH is set
	if val, exists := os.LookupEnv("HOTLINE_CONFIG_PATH"); exists {
		path = val
	}

	raw, err := ioutil.ReadFile(os.ExpandEnv(path))
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(raw, &config)
	if err != nil {
		panic(err)
	}

	return config
}
