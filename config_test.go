package main

import (
	"storeinspection/config"
	"testing"
)

func TestConfig(t *testing.T) {
	if config.Default().DatabasePath == "" {
		t.Fatal()
	}
	if !config.DefaultPolicy().AllowsStatus("new") {
		t.Fatal()
	}
}
