package model

import (
	"encoding/json"
	"fmt"
	"time"
)

func (r Record) Marshal() ([]byte, error)        { return json.Marshal(r) }
func UnmarshalRecord(b []byte) (Record, error)   { var r Record; e := json.Unmarshal(b, &r); return r, e }
func (e Event) Marshal() ([]byte, error)         { return json.Marshal(e) }
func UnmarshalEvent(b []byte) (Event, error)     { var e Event; x := json.Unmarshal(b, &e); return e, x }
func (a Audit) Marshal() ([]byte, error)         { return json.Marshal(a) }
func UnmarshalAudit(b []byte) (Audit, error)     { var a Audit; x := json.Unmarshal(b, &a); return a, x }
func (r Record) String() string                  { return fmt.Sprintf("%s/%s:%s", r.StoreID, r.ID, r.Status) }
func (r Record) Age(now time.Time) time.Duration { return now.Sub(r.UpdatedAt) }
func (r Record) IsCritical() bool                { return r.Severity >= 4 }
func (r Record) IsResolved() bool {
	return r.Status == "resolved" || r.Status == "closed" || r.Status == "archived"
}
