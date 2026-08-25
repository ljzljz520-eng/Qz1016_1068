package service

import (
	"context"
	"storeinspection/model"
)

func (s *Service) Cancel(ctx context.Context, id, actor string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if err := s.Proc.Cancel(ctx, id, actor); err != nil {
		return err
	}
	return s.Audit.Record("cancel", id, actor)
}
func (s *Service) Archive(ctx context.Context, id string) (model.Record, error) {
	r, e := s.Proc.Archive(ctx, id)
	if e == nil {
		_ = s.Audit.Record("archive", id, "system")
	}
	return r, e
}
func (s *Service) History(id string) ([]model.Audit, error) { return s.Audit.History(id) }
