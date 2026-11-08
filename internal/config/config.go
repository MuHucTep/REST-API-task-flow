package config

import "os"

type Config struct {
	HTTPAddr string
}

func Load() Config {
	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		addr = ":8000"
	}

	return Config{
		HTTPAddr: addr,
	}
}
