package config

import (
	"os"
	"strconv"
	"strings"
)

type Settings struct {
	Config
	Debug          bool
	Workers        int
	TimeoutSeconds int
}

func Load() Settings {
	s := Settings{Config: Default(), Workers: 4, TimeoutSeconds: 30}
	s.Debug = os.Getenv("INSPECTION_DEBUG") == "1"
	if n, e := strconv.Atoi(os.Getenv("INSPECTION_WORKERS")); e == nil && n > 0 {
		s.Workers = n
	}
	if n, e := strconv.Atoi(os.Getenv("INSPECTION_TIMEOUT")); e == nil && n > 0 {
		s.TimeoutSeconds = n
	}
	return s
}
func (s Settings) Tags() []string {
	v := os.Getenv("INSPECTION_TAGS")
	if v == "" {
		return []string{"store", "inspection"}
	}
	out := []string{}
	for _, x := range strings.Split(v, ",") {
		x = strings.TrimSpace(x)
		if x != "" {
			out = append(out, x)
		}
	}
	return out
}
func (s Settings) IsProduction() bool { return os.Getenv("INSPECTION_ENV") == "production" }
func (s Settings) EffectiveWorkers() int {
	if s.Workers < 1 {
		return 1
	}
	if s.Workers > 64 {
		return 64
	}
	return s.Workers
}
