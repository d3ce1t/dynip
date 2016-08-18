package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	SubDomainName string `yaml:"subdomain_name"`
	DomainName    string `yaml:"domain_name"`
	APIPassword   string `yaml:"password"`
	UpdateFreqSec int    `yaml:"update_freq_sec"`
	VerifyChange  bool   `yaml:"verify_change,omitempty"`
}

func loadConfigFromFile(file string) (*Config, error) {

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	config := &Config{}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
