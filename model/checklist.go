package model

type ChecklistItem struct {
	Code     string `json:"code"`
	Label    string `json:"label"`
	Required bool   `json:"required"`
	Complete bool   `json:"complete"`
	Finding  string `json:"finding"`
}
type Checklist struct {
	RecordID string          `json:"record_id"`
	Items    []ChecklistItem `json:"items"`
}

func (c Checklist) CompleteCount() int {
	n := 0
	for _, i := range c.Items {
		if i.Complete {
			n++
		}
	}
	return n
}
func (c Checklist) RequiredComplete() bool {
	for _, i := range c.Items {
		if i.Required && !i.Complete {
			return false
		}
	}
	return true
}
func (c *Checklist) Mark(code string, complete bool, finding string) bool {
	for n := range c.Items {
		if c.Items[n].Code == code {
			c.Items[n].Complete = complete
			c.Items[n].Finding = finding
			return true
		}
	}
	return false
}
func DefaultChecklist(id string) Checklist {
	return Checklist{RecordID: id, Items: []ChecklistItem{{"safety", "Safety signage", true, false, ""}, {"stock", "Stock rotation", true, false, ""}, {"clean", "Cleanliness", false, false, ""}, {"cash", "Cash controls", true, false, ""}}}
}
