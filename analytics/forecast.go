package analytics

import (
	"sort"
	"storeinspection/model"
	"time"
)

type Forecast struct {
	Store    string
	Open     int
	Critical int
	Due      time.Time
	Score    float64
}

func ForecastStores(rows []model.Record) []Forecast {
	m := map[string][]model.Record{}
	for _, r := range rows {
		m[r.StoreID] = append(m[r.StoreID], r)
	}
	out := []Forecast{}
	for store, rs := range m {
		f := Forecast{Store: store, Due: time.Now().UTC().Add(24 * time.Hour)}
		for _, r := range rs {
			if r.IsOpen() {
				f.Open++
			}
			if r.Severity >= 4 && r.IsOpen() {
				f.Critical++
				f.Score += float64(r.Severity)
			}
		}
		if f.Open > 0 {
			f.Score /= float64(f.Open)
		}
		out = append(out, f)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Score > out[j].Score })
	return out
}
func DueRecords(rows []model.Record, now time.Time) []model.Record {
	out := []model.Record{}
	for _, r := range rows {
		if r.IsOpen() && r.UpdatedAt.Add(72*time.Hour).Before(now) {
			out = append(out, r)
		}
	}
	return out
}
func CriticalRecords(rows []model.Record) []model.Record {
	out := []model.Record{}
	for _, r := range rows {
		if r.Severity >= 4 && r.IsOpen() {
			out = append(out, r)
		}
	}
	return out
}
func ResolvedRecords(rows []model.Record) []model.Record {
	out := []model.Record{}
	for _, r := range rows {
		if r.Status == "resolved" {
			out = append(out, r)
		}
	}
	return out
}
func ClosedRecords(rows []model.Record) []model.Record {
	out := []model.Record{}
	for _, r := range rows {
		if r.Status == "closed" || r.Status == "archived" {
			out = append(out, r)
		}
	}
	return out
}
func CountStatus(rows []model.Record, status string) int {
	n := 0
	for _, r := range rows {
		if r.Status == status {
			n++
		}
	}
	return n
}
func CountStore(rows []model.Record, store string) int {
	n := 0
	for _, r := range rows {
		if r.StoreID == store {
			n++
		}
	}
	return n
}
func CountSeverity(rows []model.Record, s int) int {
	n := 0
	for _, r := range rows {
		if r.Severity == s {
			n++
		}
	}
	return n
}
func Oldest(rows []model.Record) (model.Record, bool) {
	if len(rows) == 0 {
		return model.Record{}, false
	}
	out := rows[0]
	for _, r := range rows[1:] {
		if r.UpdatedAt.Before(out.UpdatedAt) {
			out = r
		}
	}
	return out, true
}
func Newest(rows []model.Record) (model.Record, bool) {
	if len(rows) == 0 {
		return model.Record{}, false
	}
	out := rows[0]
	for _, r := range rows[1:] {
		if r.UpdatedAt.After(out.UpdatedAt) {
			out = r
		}
	}
	return out, true
}
