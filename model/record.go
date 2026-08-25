package model

import "time"

type Record struct {
	ID        string    `json:"id"`
	StoreID   string    `json:"store_id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	Severity  int       `json:"severity"`
	Assignee  string    `json:"assignee"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Version   int       `json:"version"`
}

func NewRecord(id, store, title string, severity int) Record {
	now := time.Now().UTC()
	return Record{ID: id, StoreID: store, Title: title, Severity: severity, Status: "new", CreatedAt: now, UpdatedAt: now, Version: 1}
}
func (r Record) IsOpen() bool { return r.Status != "closed" && r.Status != "archived" }
func (r *Record) ApplyStatus(status, notes string) error {
	if status == "" {
		return ErrInvalidStatus
	}
	r.Status = status
	r.Notes = notes
	r.UpdatedAt = time.Now().UTC()
	r.Version++
	return nil
}
