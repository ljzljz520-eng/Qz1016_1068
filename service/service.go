package service

import (
	"context"
	"storeinspection/audit"
	"storeinspection/config"
	"storeinspection/model"
	"storeinspection/storage"
	"storeinspection/workflow"
)

type Service struct {
	DB     *storage.DB
	Reg    *workflow.Registration
	Proc   *workflow.Processor
	Query  *workflow.Query
	Audit  *audit.Logger
	Policy config.Policy
}

func New(db *storage.DB) *Service {
	return &Service{DB: db, Reg: workflow.NewRegistration(db), Proc: workflow.NewProcessor(db), Query: workflow.NewQuery(db), Audit: audit.New(db), Policy: config.DefaultPolicy()}
}
func (s *Service) Register(ctx context.Context, r model.Record) (model.Record, error) {
	r.Severity = s.Policy.NormalizeSeverity(r.Severity)
	if r.Status == "" {
		r.Status = "new"
	}
	if err := workflow.ValidateRecord(r, s.Policy); err != nil {
		return r, err
	}
	out, err := s.Reg.Submit(ctx, r)
	if err == nil {
		_ = s.Audit.Record("register", r.ID, "system")
	}
	return out, err
}
func (s *Service) Transition(ctx context.Context, id, actor, action, notes string) (model.Record, error) {
	status := workflow.NextStatus("", action)
	if status == "" {
		return model.Record{}, model.ErrInvalidStatus
	}
	out, err := s.Proc.Process(ctx, id, actor, status, notes)
	if err == nil {
		_ = s.Audit.Record(action, id, actor)
	}
	return out, err
}
func (s *Service) Search(ctx context.Context, store, status string) ([]model.Record, error) {
	return s.Query.Find(ctx, store, status)
}
func (s *Service) Details(ctx context.Context, id string) (model.Record, []model.Event, error) {
	r, e := s.DB.GetRecord(id)
	if e != nil {
		return r, nil, e
	}
	events, e := s.Query.Timeline(ctx, id)
	return r, events, e
}
