package analytics

import "storeinspection/model"

type Catalog struct {
	Names map[string]string
	Codes map[string]int
}

func NewCatalog() Catalog {
	return Catalog{Names: map[string]string{"new": "New", "assigned": "Assigned", "in_progress": "In progress", "resolved": "Resolved", "closed": "Closed", "archived": "Archived"}, Codes: map[string]int{"low": 1, "medium": 3, "high": 5}}
}
func (c Catalog) Name(s string) string {
	if v, ok := c.Names[s]; ok {
		return v
	}
	return s
}
func (c Catalog) Code(s string) int {
	if v, ok := c.Codes[s]; ok {
		return v
	}
	return 0
}
func (c Catalog) Statuses() []string {
	return []string{"new", "assigned", "in_progress", "resolved", "closed", "archived"}
}
func (c Catalog) SeverityName(n int) string {
	if n <= 1 {
		return "low"
	}
	if n <= 3 {
		return "medium"
	}
	return "high"
}
func (c Catalog) IsKnown(s string) bool { _, ok := c.Names[s]; return ok }
func (c Catalog) Describe(r model.Record) string {
	return c.Name(r.Status) + " / " + c.SeverityName(r.Severity)
}
func CountOpenByStore(rows []model.Record) map[string]int {
	m := map[string]int{}
	for _, r := range rows {
		if r.IsOpen() {
			m[r.StoreID]++
		}
	}
	return m
}
func CountClosedByStore(rows []model.Record) map[string]int {
	m := map[string]int{}
	for _, r := range rows {
		if !r.IsOpen() {
			m[r.StoreID]++
		}
	}
	return m
}
func CountCriticalByStore(rows []model.Record) map[string]int {
	m := map[string]int{}
	for _, r := range rows {
		if r.IsCritical() {
			m[r.StoreID]++
		}
	}
	return m
}
func CountAssigned(rows []model.Record) map[string]int {
	m := map[string]int{}
	for _, r := range rows {
		if r.Assignee != "" {
			m[r.Assignee]++
		}
	}
	return m
}
func Titles(rows []model.Record) []string {
	out := []string{}
	for _, r := range rows {
		out = append(out, r.Title)
	}
	return out
}
func IDs(rows []model.Record) []string {
	out := []string{}
	for _, r := range rows {
		out = append(out, r.ID)
	}
	return out
}
func Stores(rows []model.Record) []string {
	m := map[string]bool{}
	for _, r := range rows {
		m[r.StoreID] = true
	}
	out := []string{}
	for s := range m {
		out = append(out, s)
	}
	return out
}
func Statuses(rows []model.Record) []string {
	m := map[string]bool{}
	for _, r := range rows {
		m[r.Status] = true
	}
	out := []string{}
	for s := range m {
		out = append(out, s)
	}
	return out
}
func Assignments(rows []model.Record) map[string][]model.Record {
	m := map[string][]model.Record{}
	for _, r := range rows {
		m[r.Assignee] = append(m[r.Assignee], r)
	}
	return m
}
func ByID(rows []model.Record) map[string]model.Record {
	m := map[string]model.Record{}
	for _, r := range rows {
		m[r.ID] = r
	}
	return m
}
func Clone(rows []model.Record) []model.Record {
	out := make([]model.Record, len(rows))
	copy(out, rows)
	return out
}
func Limit(rows []model.Record, n int) []model.Record {
	if n < 0 {
		n = 0
	}
	if n > len(rows) {
		n = len(rows)
	}
	return rows[:n]
}
func Skip(rows []model.Record, n int) []model.Record {
	if n < 0 {
		n = 0
	}
	if n > len(rows) {
		n = len(rows)
	}
	return rows[n:]
}
func HasStatus(rows []model.Record, s string) bool {
	for _, r := range rows {
		if r.Status == s {
			return true
		}
	}
	return false
}
func HasStore(rows []model.Record, s string) bool {
	for _, r := range rows {
		if r.StoreID == s {
			return true
		}
	}
	return false
}
func HasSeverity(rows []model.Record, n int) bool {
	for _, r := range rows {
		if r.Severity == n {
			return true
		}
	}
	return false
}
func CountOpen(rows []model.Record) int {
	return CountStatus(rows, "new") + CountStatus(rows, "assigned") + CountStatus(rows, "in_progress") + CountStatus(rows, "resolved")
}
func CountArchived(rows []model.Record) int       { return CountStatus(rows, "archived") }
func CountNew(rows []model.Record) int            { return CountStatus(rows, "new") }
func CountResolved(rows []model.Record) int       { return CountStatus(rows, "resolved") }
func CountInProgress(rows []model.Record) int     { return CountStatus(rows, "in_progress") }
func CountAssignedStatus(rows []model.Record) int { return CountStatus(rows, "assigned") }
