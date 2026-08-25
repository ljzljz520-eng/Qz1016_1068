package main

import (
	"context"
	"storeinspection/model"
	"storeinspection/service"
	"storeinspection/storage"
	"testing"
)

func TestLifecycle(t *testing.T) {
	d, _ := storage.Open(t.TempDir() + "/l")
	defer d.Close()
	s := service.New(d)
	r, _ := s.Register(context.Background(), model.NewRecord("l", "21", "x", 1))
	if e := s.Cancel(context.Background(), r.ID, "u"); e != nil {
		t.Fatal(e)
	}
}
