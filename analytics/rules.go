package analytics

import (
	"sort"
	"storeinspection/model"
	"strings"
	"time"
)

type Bucket struct {
	Key      string
	Count    int
	Severity int
	Last     time.Time
}
type Snapshot struct {
	Total      int
	Open       int
	Closed     int
	ByStatus   map[string]int
	ByStore    map[string]int
	BySeverity map[int]int
	Buckets    []Bucket
}

func NewSnapshot() Snapshot {
	return Snapshot{ByStatus: map[string]int{}, ByStore: map[string]int{}, BySeverity: map[int]int{}}
}
func Build(rows []model.Record) Snapshot {
	s := NewSnapshot()
	for _, r := range rows {
		s.Total++
		s.ByStatus[r.Status]++
		s.ByStore[r.StoreID]++
		s.BySeverity[r.Severity]++
		if r.IsOpen() {
			s.Open++
		} else {
			s.Closed++
		}
	}
	s.Buckets = RankStores(s.ByStore)
	return s
}
func RankStores(in map[string]int) []Bucket {
	out := make([]Bucket, 0, len(in))
	for k, v := range in {
		out = append(out, Bucket{Key: k, Count: v})
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Count == out[j].Count {
			return out[i].Key < out[j].Key
		}
		return out[i].Count > out[j].Count
	})
	return out
}
func Search(rows []model.Record, q string) []model.Record {
	q = strings.ToLower(strings.TrimSpace(q))
	if q == "" {
		return append([]model.Record{}, rows...)
	}
	out := []model.Record{}
	for _, r := range rows {
		if strings.Contains(strings.ToLower(r.Title), q) || strings.Contains(strings.ToLower(r.Notes), q) || strings.Contains(strings.ToLower(r.StoreID), q) {
			out = append(out, r)
		}
	}
	return out
}
func GroupByStatus(rows []model.Record) map[string][]model.Record {
	out := map[string][]model.Record{}
	for _, r := range rows {
		out[r.Status] = append(out[r.Status], r)
	}
	return out
}
func GroupByStore(rows []model.Record) map[string][]model.Record {
	out := map[string][]model.Record{}
	for _, r := range rows {
		out[r.StoreID] = append(out[r.StoreID], r)
	}
	return out
}
func AverageSeverity(rows []model.Record) float64 {
	if len(rows) == 0 {
		return 0
	}
	sum := 0
	for _, r := range rows {
		sum += r.Severity
	}
	return float64(sum) / float64(len(rows))
}
func HighestSeverity(rows []model.Record) int {
	max := 0
	for _, r := range rows {
		if r.Severity > max {
			max = r.Severity
		}
	}
	return max
}
func OpenRatio(rows []model.Record) float64 {
	if len(rows) == 0 {
		return 0
	}
	n := 0
	for _, r := range rows {
		if r.IsOpen() {
			n++
		}
	}
	return float64(n) / float64(len(rows))
}
func StatusOrder() []string {
	return []string{"new", "assigned", "in_progress", "resolved", "closed", "archived"}
}
func IsTerminal(s string) bool        { return s == "closed" || s == "archived" }
func IsEscalated(r model.Record) bool { return r.Severity >= 4 && r.IsOpen() }
func NeedsReview(r model.Record) bool { return r.Status == "resolved" || r.Status == "closed" }
