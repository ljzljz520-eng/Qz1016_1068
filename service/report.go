package service

import (
	"context"
	"storeinspection/model"
)

type Report struct {
	Status map[string]int `json:"status"`
	Total  int            `json:"total"`
	Open   int            `json:"open"`
}

func (s *Service) Report(ctx context.Context, store string) (Report, error) {
	m, e := s.Query.Summary(ctx, store)
	if e != nil {
		return Report{}, e
	}
	rows, e := s.Search(ctx, store, "")
	if e != nil {
		return Report{}, e
	}
	r := Report{Status: m, Total: len(rows)}
	for _, v := range rows {
		if v.IsOpen() {
			r.Open++
		}
	}
	return r, nil
}
func (s *Service) SeedUser(u model.User) error { return s.DB.PutUser(u) }
