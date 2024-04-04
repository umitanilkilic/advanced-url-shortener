package config

import (
	"os"
	"strings"
)

// TODO: Improve config package

// Config is a map of environment variables
var Config *map[string]string

// LoadConfig loads the environment variables into the Config map
func init() {
	cfg := make(map[string]string)
	for _, e := range os.Environ() {
		if i := strings.Index(e, "="); i >= 0 {
			cfg[e[:i]] = e[i+1:]
		}
	}
	Config = &cfg
}
