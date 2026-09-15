package examples_test

import (
	"embed"
	"testing"

	"github.com/joaomagfreitas/configx"
)

//go:embed *.yaml
var fs embed.FS

func TestUnmarshallSqlite(t *testing.T) {
	cfg, err := configx.UnmarshalFs[configx.Sqlite](fs, "sqlite.yaml")
	if err != nil {
		t.Fatal(err)
	}

	exp := configx.Sqlite{
		Path: "file:///tmp/foo.db",
	}

	if cfg != exp {
		t.Fatalf("expected %v, got: %v", exp, cfg)
	}
}
