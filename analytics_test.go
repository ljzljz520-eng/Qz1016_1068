package main

import (
	"storeinspection/analytics"
	"storeinspection/model"
	"testing"
)

func TestAnalytics(t *testing.T) {
	rows := []model.Record{model.NewRecord("1", "21", "a", 4), model.NewRecord("2", "22", "b", 1)}
	s := analytics.Build(rows)
	if s.Total != 2 || s.Open != 2 {
		t.Fatal(s)
	}
}
