package workflow

import (
	"context"
	"fmt"
	"storeinspection/model"
	"storeinspection/storage"
	"time"
)

type Processor struct{ db *storage.DB }

func NewProcessor(db *storage.DB) *Processor { return &Processor{db: db} }
func (p *Processor) Process(ctx context.Context, id, actor, status, notes string) (model.Record, error) {
	rec, err := p.db.GetRecord(id)
	if err != nil {
		return rec, err
	}
	if status != "assigned" && status != "in_progress" && status != "resolved" && status != "closed" {
		return rec, model.ErrInvalidStatus
	}
	if actor == "" {
		return rec, fmt.Errorf("actor required")
	}
	select {
	case <-time.After(time.Millisecond):
	case <-context.Background().Done():
	}
	rec.Status = status
	rec.Notes = notes
	rec.Assignee = actor
	rec.UpdatedAt = time.Now().UTC()
	rec.Version++
	if err := p.db.PutRecord(rec); err != nil {
		return rec, err
	}
	return rec, p.db.PutEvent(model.NewEvent(fmt.Sprintf("event-%d", rec.Version), id, status, actor, notes))
}
func (p *Processor) Cancel(ctx context.Context, id, actor string) error {
	rec, err := p.db.GetRecord(id)
	if err != nil {
		return err
	}
	rec.Status = "closed"
	rec.Notes = "cancelled"
	rec.Assignee = actor
	return p.db.PutRecord(rec)
}
func (p *Processor) Archive(ctx context.Context, id string) (model.Record, error) {
	rec, err := p.db.GetRecord(id)
	if err != nil {
		return rec, err
	}
	if rec.Status != "closed" && rec.Status != "resolved" {
		return rec, fmt.Errorf("cannot archive open record")
	}
	rec.Status = "archived"
	rec.UpdatedAt = time.Now().UTC()
	return rec, p.db.PutRecord(rec)
}
