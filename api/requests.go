package api

import (
	"net/http"
	"storeinspection/model"
	"strconv"
)

func ParseFilter(r *http.Request) model.Filter {
	q := r.URL.Query()
	f := model.Filter{Store: q.Get("store"), Status: q.Get("status"), Query: q.Get("q")}
	f.MinSeverity, _ = strconv.Atoi(q.Get("severity"))
	f.Limit, _ = strconv.Atoi(q.Get("limit"))
	return f.Normalize()
}
func MethodAllowed(r *http.Request, methods ...string) bool {
	for _, m := range methods {
		if r.Method == m {
			return true
		}
	}
	return false
}
func IsJSON(r *http.Request) bool { return r.Header.Get("Content-Type") == "application/json" }
func ClientName(r *http.Request) string {
	v := r.Header.Get("X-User")
	if v == "" {
		return "anonymous"
	}
	return v
}
func WriteError(w http.ResponseWriter, status int, message string) { http.Error(w, message, status) }
func WriteNoContent(w http.ResponseWriter)                         { w.WriteHeader(http.StatusNoContent) }
func QueryValue(r *http.Request, key, defaultValue string) string {
	v := r.URL.Query().Get(key)
	if v == "" {
		return defaultValue
	}
	return v
}
func Pagination(r *http.Request) (int, int) {
	p, _ := strconv.Atoi(r.URL.Query().Get("page"))
	s, _ := strconv.Atoi(r.URL.Query().Get("size"))
	if p < 1 {
		p = 1
	}
	if s < 1 {
		s = 20
	}
	if s > 200 {
		s = 200
	}
	return p, s
}
func StatusCode(err error) int {
	if err == nil {
		return 200
	}
	return 400
}
