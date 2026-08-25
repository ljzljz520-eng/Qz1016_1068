package analytics

import "storeinspection/model"

func Label(r model.Record) string {
	if r.IsCritical() {
		return "critical"
	}
	if r.IsResolved() {
		return "complete"
	}
	return "standard"
}
func Labels(rows []model.Record) map[string]string {
	out := map[string]string{}
	for _, r := range rows {
		out[r.ID] = Label(r)
	}
	return out
}
func CriticalRatio(rows []model.Record) float64 {
	if len(rows) == 0 {
		return 0
	}
	return float64(len(CriticalRecords(rows))) / float64(len(rows))
}
func CompleteRatio(rows []model.Record) float64 {
	if len(rows) == 0 {
		return 0
	}
	return float64(len(ClosedRecords(rows))) / float64(len(rows))
}
func StoreRatio(rows []model.Record, store string) float64 {
	if len(rows) == 0 {
		return 0
	}
	return float64(CountStore(rows, store)) / float64(len(rows))
}
func SeverityRatio(rows []model.Record, s int) float64 {
	if len(rows) == 0 {
		return 0
	}
	return float64(CountSeverity(rows, s)) / float64(len(rows))
}
func StatusRatio(rows []model.Record, s string) float64 {
	if len(rows) == 0 {
		return 0
	}
	return float64(CountStatus(rows, s)) / float64(len(rows))
}
func OpenBySeverity(rows []model.Record) map[int]int {
	out := map[int]int{}
	for _, r := range rows {
		if r.IsOpen() {
			out[r.Severity]++
		}
	}
	return out
}
func ClosedBySeverity(rows []model.Record) map[int]int {
	out := map[int]int{}
	for _, r := range rows {
		if !r.IsOpen() {
			out[r.Severity]++
		}
	}
	return out
}
func ActiveByAssignee(rows []model.Record) map[string]int {
	out := map[string]int{}
	for _, r := range rows {
		if r.IsOpen() {
			out[r.Assignee]++
		}
	}
	return out
}
func SeverityLabels(rows []model.Record) map[string]string {
	out := map[string]string{}
	for _, r := range rows {
		out[r.ID] = NewCatalog().SeverityName(r.Severity)
	}
	return out
}
func HasCritical(rows []model.Record) bool { return len(CriticalRecords(rows)) > 0 }
func HasOpen(rows []model.Record) bool { return len(Active(rows)) > 0 }
func HasClosed(rows []model.Record) bool { return len(ClosedRecords(rows)) > 0 }
func HasUrgent(rows []model.Record) bool { return len(Urgent(rows)) > 0 }
func StoreNames(rows []model.Record) []string { return Stores(rows) }
func StatusNames(rows []model.Record) []string { return Statuses(rows) }
func SeverityNames(rows []model.Record) []string { out:=[]string{}; c:=NewCatalog(); for _,r:=range rows{out=append(out,c.SeverityName(r.Severity))}; return out }
func RecordLabels(rows []model.Record) []string { out:=[]string{}; for _,r:=range rows{out=append(out,Label(r))}; return out }
func OpenIDs(rows []model.Record) []string { return IDs(Active(rows)) }
func CriticalIDs(rows []model.Record) []string { return IDs(CriticalRecords(rows)) }
func ClosedIDs(rows []model.Record) []string { return IDs(ClosedRecords(rows)) }
func StoreIDs(rows []model.Record,store string) []string { return IDs(ForStore(rows,store)) }
func StatusIDs(rows []model.Record,status string) []string { return IDs(ForStatus(rows,status)) }
