package config

import (
	"io/ioutil"
	"log"
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
		log.Fatalln("failed to load config:", err)
		return nil
	}

	err = yaml.Unmarshal(raw, &config)
	if err != nil {
		log.Panicln(err)
	}

	return config
}
