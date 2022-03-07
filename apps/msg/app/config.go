package app

import (
	"log"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"

	registryv1 "github.com/yusank/goim/api/config/registry/v1"
	configv1 "github.com/yusank/goim/api/config/v1"
)

type Config struct {
	*configv1.Service `json:",inline"`
	config.Config     `json:"-"`
	Filepath          string
}

func NewConfig() *Config {
	return &Config{
		Service: new(configv1.Service),
	}
}

type Registry struct {
	*registryv1.Registry `json:",inline"`
	config.Config        `json:"-"`
	FilePath             string
}

func NewRegistry() *Registry {
	return &Registry{
		Registry: new(registryv1.Registry),
	}
}

func ParseConfig(fp string) (*Config, *Registry) {
	c := config.New(
		config.WithSource(
			file.NewSource(fp),
		),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}

	cfg := NewConfig()
	cfg.Config = c
	// Unmarshal the config to struct
	if err := c.Scan(cfg); err != nil {
		panic(err)
	}
	log.Printf("%+v", cfg)

	reg := NewRegistry()
	reg.Config = c
	if err := c.Scan(reg); err != nil {
		panic(err)
	}
	log.Printf("%+v", reg)
	reg.Name = cfg.GetName()
	return cfg, reg
}
