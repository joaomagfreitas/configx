package configx

import (
	"os"

	"github.com/goccy/go-yaml"
)

func Unmarshal[Config any](path string) (Config, error) {
	var cfg Config
	err := unmarshal(path, &cfg)

	return cfg, err
}

func unmarshal(p string, cfg interface{}) error {
	var bs []byte

	bs, err := os.ReadFile(p)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(bs, &cfg)
}
