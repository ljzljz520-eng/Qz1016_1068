package main

import (
	"storeinspection/model"
	"testing"
)

func TestRecordModel(t *testing.T) {
	r := model.NewRecord("a", "21", "x", 9)
	if r.Severity != 9 || !r.IsOpen() {
		t.Fatal()
	}
	if r.ApplyStatus("closed", "") != nil {
		t.Fatal()
	}
}
func TestChecklist(t *testing.T) {
	c := model.DefaultChecklist("x")
	if c.RequiredComplete() {
		t.Fatal()
	}
	c.Mark("safety", true, "")
	c.Mark("stock", true, "")
	c.Mark("cash", true, "")
	if !c.RequiredComplete() {
		t.Fatal()
	}
}
