package main

import (
	"context"
	"storeinspection/model"
	"storeinspection/service"
	"storeinspection/storage"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := t.TempDir() + "/p.db"
	db, e := storage.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	s := service.New(db)
	if _, e = s.Register(context.Background(), model.NewRecord("persist", "21", "saved", 1)); e != nil {
		t.Fatal(e)
	}
	db.Close()
	db, e = storage.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer db.Close()
	if _, e = db.GetRecord("persist"); e != nil {
		t.Fatal(e)
	}
}
