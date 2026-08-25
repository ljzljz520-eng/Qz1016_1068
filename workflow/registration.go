package workflow

import (
	"context"
	"fmt"
	"storeinspection/model"
	"storeinspection/storage"
)

type Registration struct{ db *storage.DB }

func NewRegistration(db *storage.DB) *Registration { return &Registration{db: db} }
func (r *Registration) Submit(ctx context.Context, rec model.Record) (model.Record, error) {
	if err := ctx.Err(); err != nil {
		return rec, err
	}
	if rec.ID == "" || rec.StoreID == "" || rec.Title == "" {
		return rec, fmt.Errorf("missing registration data")
	}
	if rec.Severity < 1 {
		rec.Severity = 1
	}
	if rec.Status == "" {
		rec.Status = "new"
	}
	if err := r.db.PutRecord(rec); err != nil {
		return rec, err
	}
	return rec, nil
}
func (r *Registration) Assign(ctx context.Context, id, user string) (model.Record, error) {
	rec, err := r.db.GetRecord(id)
	if err != nil {
		return rec, err
	}
	if user == "" {
		return rec, fmt.Errorf("assignee required")
	}
	rec.Assignee = user
	rec.Status = "assigned"
	rec, err = r.save(ctx, rec, "assign")
	return rec, err
}
func (r *Registration) save(ctx context.Context, rec model.Record, kind string) (model.Record, error) {
	if ctx.Err() != nil {
		return rec, ctx.Err()
	}
	if err := rec.ApplyStatus(rec.Status, rec.Notes); err != nil {
		return rec, err
	}
	e := model.NewEvent(fmt.Sprintf("%s-%d", kind, rec.Version), rec.ID, kind, rec.Assignee, rec.Notes)
	if err := r.db.PutRecord(rec); err != nil {
		return rec, err
	}
	return rec, r.db.PutEvent(e)
}
