package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type MetaConfig struct {
	Port     int32  `yaml:"port"`
	RunMode  string `yaml:"run-mode"`
	LogPath  string `yaml:"log-path"`
	LogName  string `yaml:"log-name"`
	DefLimit int32  `yaml:"def-limit"`
	MaxLimit int32  `yaml:"max-limit"`
}

type Config struct {
	Meta *MetaConfig `yaml:"meta"`
}

func Load(path string) (*Config, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = yaml.Unmarshal(f, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
