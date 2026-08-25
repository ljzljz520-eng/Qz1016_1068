package model

import "time"

type Event struct {
	ID       string    `json:"id"`
	RecordID string    `json:"record_id"`
	Kind     string    `json:"kind"`
	Actor    string    `json:"actor"`
	Detail   string    `json:"detail"`
	At       time.Time `json:"at"`
}
type Audit struct {
	ID       string    `json:"id"`
	Action   string    `json:"action"`
	RecordID string    `json:"record_id"`
	Actor    string    `json:"actor"`
	At       time.Time `json:"at"`
}

func NewEvent(id, record, kind, actor, detail string) Event {
	return Event{ID: id, RecordID: record, Kind: kind, Actor: actor, Detail: detail, At: time.Now().UTC()}
}
func NewAudit(id, action, record, actor string) Audit {
	return Audit{ID: id, Action: action, RecordID: record, Actor: actor, At: time.Now().UTC()}
}
