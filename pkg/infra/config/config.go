package config

import (
	"github.com/jinzhu/configor"
)

type Config struct {
	Name   string `default:"Cadastro Power"`
	Server struct {
		Port string `default:"8080"`
	}
	Database struct {
		Cadastro struct {
			Host     string `default:"localhost"`
			Port     int    `default:"3306"`
			UserName string `default:"user"`
			Password string `default:"pass"`
			Database string `default:"cadastropower-db"`
		}
	}
}

func NewConfig() (*Config, error) {

	options := &configor.Config{Verbose: true, Debug: true}

	config := &Config{}
	if err := configor.New(options).Load(config, "config.yml", "pkg/infra/config/config.yml"); err != nil {
		return nil, err
	}

	return config, nil
}
