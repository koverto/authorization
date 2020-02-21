package handler

import (
	"fmt"

	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source"
)

type Config struct {
	MongoUrl string `json:"mongourl"`
	Name     string
}

func NewConfig(name string, source ...source.Source) (*Config, error) {
	if err := config.Load(source...); err != nil {
		return nil, err
	}

	conf := &Config{
		MongoUrl: "mongodb://localhost:27017",
		Name:     name,
	}

	err := config.Scan(conf)
	return conf, err
}

func (c *Config) ID() string {
	return fmt.Sprintf("com.koverto.svc.%s", c.Name)
}
