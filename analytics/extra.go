package analytics

import (
	"storeinspection/model"
	"time"
)

func Active(rows []model.Record) []model.Record {
	return filter(rows, func(r model.Record) bool { return r.IsOpen() })
}
func Urgent(rows []model.Record) []model.Record {
	return filter(rows, func(r model.Record) bool { return r.IsCritical() && r.IsOpen() })
}
func WithAssignee(rows []model.Record) []model.Record {
	return filter(rows, func(r model.Record) bool { return r.Assignee != "" })
}
func WithoutAssignee(rows []model.Record) []model.Record {
	return filter(rows, func(r model.Record) bool { return r.Assignee == "" })
}
func ForStore(rows []model.Record, s string) []model.Record {
	return filter(rows, func(r model.Record) bool { return r.StoreID == s })
}
func ForStatus(rows []model.Record, s string) []model.Record {
	return filter(rows, func(r model.Record) bool { return r.Status == s })
}
func Before(rows []model.Record, t time.Time) []model.Record {
	return filter(rows, func(r model.Record) bool { return r.UpdatedAt.Before(t) })
}
func After(rows []model.Record, t time.Time) []model.Record {
	return filter(rows, func(r model.Record) bool { return r.UpdatedAt.After(t) })
}
func filter(rows []model.Record, fn func(model.Record) bool) []model.Record {
	out := []model.Record{}
	for _, r := range rows {
		if fn(r) {
			out = append(out, r)
		}
	}
	return out
}
func MapTitles(rows []model.Record) map[string]string {
	m := map[string]string{}
	for _, r := range rows {
		m[r.ID] = r.Title
	}
	return m
}
func MapStatuses(rows []model.Record) map[string]string {
	m := map[string]string{}
	for _, r := range rows {
		m[r.ID] = r.Status
	}
	return m
}
func MapSeverity(rows []model.Record) map[string]int {
	m := map[string]int{}
	for _, r := range rows {
		m[r.ID] = r.Severity
	}
	return m
}
func MapStores(rows []model.Record) map[string]string {
	m := map[string]string{}
	for _, r := range rows {
		m[r.ID] = r.StoreID
	}
	return m
}
func MapAssignees(rows []model.Record) map[string]string {
	m := map[string]string{}
	for _, r := range rows {
		m[r.ID] = r.Assignee
	}
	return m
}
func TotalSeverity(rows []model.Record) int {
	n := 0
	for _, r := range rows {
		n += r.Severity
	}
	return n
}
func TotalVersion(rows []model.Record) int {
	n := 0
	for _, r := range rows {
		n += r.Version
	}
	return n
}
func MaxVersion(rows []model.Record) int {
	n := 0
	for _, r := range rows {
		if r.Version > n {
			n = r.Version
		}
	}
	return n
}
func MinSeverity(rows []model.Record) int {
	if len(rows) == 0 {
		return 0
	}
	n := rows[0].Severity
	for _, r := range rows[1:] {
		if r.Severity < n {
			n = r.Severity
		}
	}
	return n
}
func First(rows []model.Record) (model.Record, bool) {
	if len(rows) == 0 {
		return model.Record{}, false
	}
	return rows[0], true
}
func Last(rows []model.Record) (model.Record, bool) {
	if len(rows) == 0 {
		return model.Record{}, false
	}
	return rows[len(rows)-1], true
}
func Any(rows []model.Record, fn func(model.Record) bool) bool {
	for _, r := range rows {
		if fn(r) {
			return true
		}
	}
	return false
}
func All(rows []model.Record, fn func(model.Record) bool) bool {
	for _, r := range rows {
		if !fn(r) {
			return false
		}
	}
	return true
}
func None(rows []model.Record, fn func(model.Record) bool) bool { return !Any(rows, fn) }
func Transform(rows []model.Record, fn func(model.Record) model.Record) []model.Record {
	out := make([]model.Record, 0, len(rows))
	for _, r := range rows {
		out = append(out, fn(r))
	}
	return out
}
func CountIf(rows []model.Record, fn func(model.Record) bool) int {
	n := 0
	for _, r := range rows {
		if fn(r) {
			n++
		}
	}
	return n
}
func Touch(rows []model.Record, t time.Time) []model.Record {
	return Transform(rows, func(r model.Record) model.Record { r.UpdatedAt = t; return r })
}
func Escalate(rows []model.Record) []model.Record {
	return Transform(rows, func(r model.Record) model.Record {
		if r.Severity < 5 {
			r.Severity++
		}
		return r
	})
}
func Normalize(rows []model.Record) []model.Record {
	return Transform(rows, func(r model.Record) model.Record {
		if r.Status == "" {
			r.Status = "new"
		}
		if r.Version < 1 {
			r.Version = 1
		}
		return r
	})
}
