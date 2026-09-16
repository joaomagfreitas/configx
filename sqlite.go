package configx

import (
	"fmt"
)

type Sqlite struct {
	Path string
}

func (cfg Sqlite) Conn() string {
	return fmt.Sprintf("file:%s", cfg.Path)
}
