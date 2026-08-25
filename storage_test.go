package main

import (
	"storeinspection/model"
	"storeinspection/storage"
	"testing"
)

func TestStorageRoundTrip(t *testing.T) {
	d, e := storage.Open(t.TempDir() + "/s")
	if e != nil {
		t.Fatal(e)
	}
	defer d.Close()
	r := model.NewRecord("s", "21", "x", 1)
	if e = d.PutRecord(r); e != nil {
		t.Fatal(e)
	}
	if _, e = d.GetRecord("s"); e != nil {
		t.Fatal(e)
	}
}
