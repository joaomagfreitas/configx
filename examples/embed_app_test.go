package examples_test

import (
	"testing"

	"github.com/joaomagfreitas/configx"
)

func TestUnmarshallEmbedApp(t *testing.T) {
	cfg, err := configx.UnmarshalFs[configx.EmbedApp](fs, "embed_app.yaml")
	if err != nil {
		t.Fatal(err)
	}

	exp := configx.EmbedApp{
		Database: configx.Sqlite{
			Path: "file:///tmp/foo.db",
		},
	}

	if cfg != exp {
		t.Fatalf("expected %v, got: %v", exp, cfg)
	}
}
