package workflow

import (
	"context"
	"storeinspection/model"
	"storeinspection/storage"
)

type Metrics struct{ db *storage.DB }

func NewMetrics(db *storage.DB) *Metrics { return &Metrics{db: db} }
func (m *Metrics) BySeverity(ctx context.Context, store string) (map[int]int, error) {
	rows, e := m.db.ListRecords(store, "")
	if e != nil {
		return nil, e
	}
	out := map[int]int{}
	for _, r := range rows {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		out[r.Severity]++
	}
	return out, nil
}
func (m *Metrics) Aging(ctx context.Context, store string) (map[string]int, error) {
	rows, e := m.db.ListRecords(store, "")
	if e != nil {
		return nil, e
	}
	out := map[string]int{"fresh": 0, "aging": 0, "stale": 0}
	for _, r := range rows {
		age := r.UpdatedAt.Sub(r.CreatedAt).Hours()
		if age < 24 {
			out["fresh"]++
		} else if age < 72 {
			out["aging"]++
		} else {
			out["stale"]++
		}
	}
	return out, nil
}
func Classify(r model.Record) string {
	if r.Status == "closed" || r.Status == "archived" {
		return "done"
	}
	if r.Severity >= 4 {
		return "urgent"
	}
	return "active"
}
