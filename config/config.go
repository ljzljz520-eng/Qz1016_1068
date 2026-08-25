package config

import "os"

type Config struct {
	DatabasePath  string
	ListenAddress string
}

func Default() Config {
	p := os.Getenv("INSPECTION_DB")
	if p == "" {
		p = "inspection.db"
	}
	a := os.Getenv("INSPECTION_ADDR")
	if a == "" {
		a = ":8080"
	}
	return Config{p, a}
}
func (c Config) Validate() error {
	if c.DatabasePath == "" || c.ListenAddress == "" {
		return os.ErrInvalid
	}
	return nil
}
