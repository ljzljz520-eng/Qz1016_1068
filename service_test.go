package main

import (
	"context"
	"storeinspection/model"
	"storeinspection/service"
	"storeinspection/storage"
	"testing"
)

func TestRecordFlow21(t *testing.T) {
	db, e := storage.Open(t.TempDir() + "/x.db")
	if e != nil {
		t.Fatal(e)
	}
	defer db.Close()
	s := service.New(db)
	r := model.NewRecord("21-1", "21", "巡检", 3)
	if e = r.ApplyStatus("resolved", "new state"); e != nil {
		t.Fatal(e)
	}
	r, e = s.Register(context.Background(), r)
	if e != nil {
		t.Fatal(e)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, e = s.Transition(ctx, r.ID, "u", "assign", "stale request")
	if e != nil {
		t.Fatalf("unexpected transition error: %v", e)
	}
	visible, _, readErr := s.Details(context.Background(), r.ID)
	if readErr != nil {
		t.Fatal(readErr)
	}
	if visible.Status != "resolved" {
		t.Fatalf("status = %q, want resolved after canceled transition", visible.Status)
	}
}
func TestWorkflowOne(t *testing.T) {
	db, e := storage.Open(t.TempDir() + "/a")
	if e != nil {
		t.Fatal(e)
	}
	defer db.Close()
	s := service.New(db)
	record := model.NewRecord("w1", "21", "entry", 2)
	if e = record.ApplyStatus("in_progress", "submitted state"); e != nil {
		t.Fatal(e)
	}
	if _, e = s.Register(context.Background(), record); e != nil {
		t.Fatal(e)
	}
	visible, _, e := s.Details(context.Background(), record.ID)
	if e != nil {
		t.Fatal(e)
	}
	if visible.Status != "in_progress" {
		t.Fatalf("status = %q, want in_progress", visible.Status)
	}
}
