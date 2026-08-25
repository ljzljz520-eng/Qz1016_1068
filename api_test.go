package main

import (
	"net/http/httptest"
	"storeinspection/api"
	"storeinspection/service"
	"storeinspection/storage"
	"testing"
)

func TestHealth(t *testing.T) {
	d, _ := storage.Open(t.TempDir() + "/a")
	defer d.Close()
	r := httptest.NewRecorder()
	api.New(service.New(d)).ServeHTTP(r, httptest.NewRequest("GET", "/health", nil))
	if r.Code != 200 {
		t.Fatal(r.Code)
	}
}
