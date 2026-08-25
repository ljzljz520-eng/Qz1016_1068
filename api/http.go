package api

import (
	"encoding/json"
	"net/http"
	"storeinspection/model"
	"storeinspection/service"
)

type Handler struct{ svc *service.Service }

func New(s *service.Service) http.Handler {
	h := &Handler{svc: s}
	m := http.NewServeMux()
	m.HandleFunc("/health", h.health)
	m.HandleFunc("/records", h.records)
	m.HandleFunc("/report", h.report)
	return m
}
func (h *Handler) health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
func (h *Handler) records(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if r.Method == http.MethodPost {
		var rec model.Record
		if e := json.NewDecoder(r.Body).Decode(&rec); e != nil {
			http.Error(w, e.Error(), 400)
			return
		}
		out, e := h.svc.Register(ctx, rec)
		if e != nil {
			http.Error(w, e.Error(), 400)
			return
		}
		write(w, out)
		return
	}
	rows, e := h.svc.Search(ctx, r.URL.Query().Get("store"), r.URL.Query().Get("status"))
	if e != nil {
		http.Error(w, e.Error(), 500)
		return
	}
	write(w, rows)
}
func (h *Handler) report(w http.ResponseWriter, r *http.Request) {
	out, e := h.svc.Report(r.Context(), r.URL.Query().Get("store"))
	if e != nil {
		http.Error(w, e.Error(), 500)
		return
	}
	write(w, out)
}
func write(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
