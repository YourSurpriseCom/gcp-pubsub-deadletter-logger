package config

import (
	log "github.com/jlentink/yaglogger"

	"github.com/vrischmann/envconfig"
)

var (
	AppConfig *Configuration
)

func init() {
	var err error
	AppConfig, err = NewConfiguration()
	if err != nil {
		log.Fatalf("Invalid environment variable configuration: %v", err)
	}
}

type Configuration struct {
	Port            int    `envconfig:"default=8888"`
	LogLevel        string `envconfig:"default=info"`
	GoogleProjectId string
}

func NewConfiguration() (*Configuration, error) {
	configuration := Configuration{}

	err := envconfig.Init(&configuration)
	return &configuration, err
}
