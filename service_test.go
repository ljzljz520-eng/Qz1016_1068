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
	r, e := s.Register(context.Background(), model.NewRecord("21-1", "21", "巡检", 3))
	if e != nil {
		t.Fatal(e)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, e = s.Transition(ctx, r.ID, "u", "resolve", "done")
	if e == nil {
		t.Fatal("expected cancellation")
	}
}
func TestWorkflowOne(t *testing.T) {
	db, e := storage.Open(t.TempDir() + "/a")
	if e != nil {
		t.Fatal(e)
	}
	defer db.Close()
	s := service.New(db)
	if _, e = s.Register(context.Background(), model.NewRecord("w1", "21", "entry", 2)); e != nil {
		t.Fatal(e)
	}
}
