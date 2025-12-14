package fixtures

import (
	"os"
	"time"

	"go-api-tests/http"

	"github.com/Nikita-Filonov/axiom"
	"gopkg.in/yaml.v3"
)

type Config struct {
	HTTP http.ClientConfig `yaml:"http"`
}

func SetConfigFixture(_ *axiom.Config) (any, func(), error) {
	out := &Config{
		HTTP: http.ClientConfig{
			URL:     "https://dummyjson.com",
			Timeout: 10 * time.Second,
		},
	}
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return out, nil, nil
	}

	if err = yaml.Unmarshal(data, out); err != nil {
		return nil, nil, err
	}

	return out, nil, nil
}

func GetConfigFixture(cfg *axiom.Config) *Config {
	return axiom.GetFixture[*Config](cfg, "config")
}
