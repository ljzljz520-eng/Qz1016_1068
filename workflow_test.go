package main

import (
	"context"
	"storeinspection/model"
	"storeinspection/service"
	"storeinspection/storage"
	"testing"
)

func TestWorkflowTwo(t *testing.T) {
	db, _ := storage.Open(t.TempDir() + "/w")
	defer db.Close()
	s := service.New(db)
	r, e := s.Register(context.Background(), model.NewRecord("w2", "21", "audit", 4))
	if e != nil {
		t.Fatal(e)
	}
	if _, e = s.Transition(context.Background(), r.ID, "m", "assign", ""); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowThree(t *testing.T) {
	db, _ := storage.Open(t.TempDir() + "/w3")
	defer db.Close()
	s := service.New(db)
	r, _ := s.Register(context.Background(), model.NewRecord("w3", "21", "track", 2))
	if _, e := s.Transition(context.Background(), r.ID, "m", "start", ""); e != nil {
		t.Fatal(e)
	}
}
