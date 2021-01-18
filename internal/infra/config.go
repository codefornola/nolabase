package infra

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	FilePath string
	Database *DbConfig `yaml:"database"`
}

type DbConfig struct {
	User string `yaml:"user"`
	Host string `yaml:"host"`
	Password string `yaml:"password"`
	Port int `yaml:"port"`
	DatabaseName string `yaml:"database_name"`
}

func (dc *DbConfig) String() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		dc.User,
		dc.Password,
		dc.Host,
		dc.Port,
		dc.DatabaseName,
	)
}

func NewConfig(filePath string) (*Config, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	yaml.Unmarshal(data, config)
	config.FilePath = filePath
	return config, nil
}


