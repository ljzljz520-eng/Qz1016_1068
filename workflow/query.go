package workflow

import (
	"context"
	"sort"
	"storeinspection/model"
	"storeinspection/storage"
)

type Query struct{ db *storage.DB }

func NewQuery(db *storage.DB) *Query { return &Query{db: db} }
func (q *Query) Find(ctx context.Context, store, status string) ([]model.Record, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	rows, err := q.db.ListRecords(store, status)
	if err != nil {
		return nil, err
	}
	sort.Slice(rows, func(i, j int) bool { return rows[i].UpdatedAt.After(rows[j].UpdatedAt) })
	return rows, nil
}
func (q *Query) Timeline(ctx context.Context, id string) ([]model.Event, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return q.db.ListEvents(id)
}
func (q *Query) Summary(ctx context.Context, store string) (map[string]int, error) {
	rows, err := q.Find(ctx, store, "")
	if err != nil {
		return nil, err
	}
	m := map[string]int{}
	for _, r := range rows {
		m[r.Status]++
	}
	return m, nil
}
