package configx

import (
	"io"
	"io/fs"
	"os"

	"github.com/goccy/go-yaml"
)

func Unmarshal[Config any](path string) (Config, error) {
	var cfg Config

	bs, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}

	err = yaml.Unmarshal(bs, &cfg)
	return cfg, err
}

func UnmarshalFs[Config any](fs fs.FS, path string) (Config, error) {
	var cfg Config

	f, err := fs.Open(path)
	if err != nil {
		return cfg, err
	}

	bs, err := io.ReadAll(f)
	if err != nil {
		return cfg, err
	}

	err = yaml.Unmarshal(bs, &cfg)
	return cfg, err
}
