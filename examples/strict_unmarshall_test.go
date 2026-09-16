package examples_test

import (
	"testing"

	"github.com/joaomagfreitas/configx"
)

func TestStrictUnmarshall(t *testing.T) {
	_, err := configx.UnmarshalFs[configx.EmbedApp](fs, "sqlite.yaml")
	if err == nil {
		t.Fatal()
	}
}
