package config

import "runtime"

type Config struct {
	PathSeparator string
}

func New() *Config {
	cfg := &Config{}

	cfg.init()

	return cfg
}

func (c *Config) init() {
	if runtime.GOOS == "windows" {
		c.PathSeparator = "\\"
	} else {
		c.PathSeparator = "/"
	}
}
