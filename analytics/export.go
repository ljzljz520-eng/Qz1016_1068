package analytics

import (
	"encoding/csv"
	"fmt"
	"io"
	"storeinspection/model"
	"strconv"
	"time"
)

func WriteCSV(w io.Writer, rows []model.Record) error {
	c := csv.NewWriter(w)
	if err := c.Write([]string{"id", "store", "title", "status", "severity", "assignee", "updated"}); err != nil {
		return err
	}
	for _, r := range rows {
		if err := c.Write([]string{r.ID, r.StoreID, r.Title, r.Status, strconv.Itoa(r.Severity), r.Assignee, r.UpdatedAt.Format(time.RFC3339)}); err != nil {
			return err
		}
	}
	c.Flush()
	return c.Error()
}
func FormatSummary(s Snapshot) string {
	return fmt.Sprintf("total=%d open=%d closed=%d", s.Total, s.Open, s.Closed)
}
func ParseSeverity(v string) int {
	n, e := strconv.Atoi(v)
	if e != nil {
		return 0
	}
	return n
}
func DateKey(t time.Time) string { return t.UTC().Format("2006-01-02") }
func DailyCounts(rows []model.Record) map[string]int {
	out := map[string]int{}
	for _, r := range rows {
		out[DateKey(r.UpdatedAt)]++
	}
	return out
}
func Since(rows []model.Record, t time.Time) []model.Record {
	out := []model.Record{}
	for _, r := range rows {
		if r.UpdatedAt.After(t) {
			out = append(out, r)
		}
	}
	return out
}
