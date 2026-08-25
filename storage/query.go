package storage

import (
	"sort"
	"storeinspection/model"
	"strings"
)

func FilterRecords(rows []model.Record, f model.Filter) []model.Record {
	f = f.Normalize()
	out := []model.Record{}
	for _, r := range rows {
		if f.Match(r) {
			out = append(out, r)
			if len(out) >= f.Limit {
				break
			}
		}
	}
	return out
}
func SortBySeverity(rows []model.Record) []model.Record {
	out := append([]model.Record{}, rows...)
	sort.SliceStable(out, func(i, j int) bool { return out[i].Severity > out[j].Severity })
	return out
}
func SortByTitle(rows []model.Record) []model.Record {
	out := append([]model.Record{}, rows...)
	sort.SliceStable(out, func(i, j int) bool { return strings.ToLower(out[i].Title) < strings.ToLower(out[j].Title) })
	return out
}
func SortByUpdated(rows []model.Record) []model.Record {
	out := append([]model.Record{}, rows...)
	sort.SliceStable(out, func(i, j int) bool { return out[i].UpdatedAt.After(out[j].UpdatedAt) })
	return out
}
func Paginate(rows []model.Record, page, size int) []model.Record {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 20
	}
	start := (page - 1) * size
	if start >= len(rows) {
		return []model.Record{}
	}
	end := start + size
	if end > len(rows) {
		end = len(rows)
	}
	return rows[start:end]
}
