package storage

import (
	"storeinspection/model"
	"strings"
)

func IndexByStore(rows []model.Record) map[string][]string {
	m := map[string][]string{}
	for _, r := range rows {
		m[r.StoreID] = append(m[r.StoreID], r.ID)
	}
	return m
}
func IndexByStatus(rows []model.Record) map[string][]string {
	m := map[string][]string{}
	for _, r := range rows {
		m[r.Status] = append(m[r.Status], r.ID)
	}
	return m
}
func IndexByAssignee(rows []model.Record) map[string][]string {
	m := map[string][]string{}
	for _, r := range rows {
		m[r.Assignee] = append(m[r.Assignee], r.ID)
	}
	return m
}
func FindID(rows []model.Record, id string) (model.Record, bool) {
	for _, r := range rows {
		if r.ID == id {
			return r, true
		}
	}
	return model.Record{}, false
}
func FindTitle(rows []model.Record, title string) []model.Record {
	out := []model.Record{}
	for _, r := range rows {
		if strings.EqualFold(r.Title, title) {
			out = append(out, r)
		}
	}
	return out
}
func FindNotes(rows []model.Record, q string) []model.Record {
	out := []model.Record{}
	for _, r := range rows {
		if strings.Contains(strings.ToLower(r.Notes), strings.ToLower(q)) {
			out = append(out, r)
		}
	}
	return out
}
func UniqueStores(rows []model.Record) []string    { return mapKeys(IndexByStore(rows)) }
func UniqueStatuses(rows []model.Record) []string  { return mapKeys(IndexByStatus(rows)) }
func UniqueAssignees(rows []model.Record) []string { return mapKeys(IndexByAssignee(rows)) }
func mapKeys(m map[string][]string) []string {
	out := []string{}
	for k := range m {
		out = append(out, k)
	}
	return out
}
