package model

import "strings"

type Filter struct {
	Store       string
	Status      string
	Query       string
	MinSeverity int
	Limit       int
}

func (f Filter) Match(r Record) bool {
	if f.Store != "" && r.StoreID != f.Store {
		return false
	}
	if f.Status != "" && r.Status != f.Status {
		return false
	}
	if f.MinSeverity > 0 && r.Severity < f.MinSeverity {
		return false
	}
	if f.Query != "" && !strings.Contains(strings.ToLower(r.Title+" "+r.Notes), strings.ToLower(f.Query)) {
		return false
	}
	return true
}
func (f Filter) Normalize() Filter {
	if f.Limit <= 0 {
		f.Limit = 100
	}
	if f.Limit > 1000 {
		f.Limit = 1000
	}
	if f.MinSeverity < 0 {
		f.MinSeverity = 0
	}
	return f
}
