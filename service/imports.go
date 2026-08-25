package service

import (
	"context"
	"fmt"
	"storeinspection/model"
)

func (s *Service) BulkRegister(ctx context.Context, records []model.Record) (int, error) {
	count := 0
	for _, r := range records {
		if err := ctx.Err(); err != nil {
			return count, err
		}
		if _, err := s.Register(ctx, r); err != nil {
			return count, fmt.Errorf("record %s: %w", r.ID, err)
		}
		count++
	}
	return count, nil
}
func (s *Service) Reassign(ctx context.Context, id, user string) (model.Record, error) {
	r, e := s.DB.GetRecord(id)
	if e != nil {
		return r, e
	}
	if user == "" {
		return r, fmt.Errorf("user required")
	}
	r.Assignee = user
	r.UpdatedAt = r.UpdatedAt.Add(0)
	e = s.DB.PutRecord(r)
	if e == nil {
		e = s.Audit.Record("reassign", id, user)
	}
	return r, e
}
