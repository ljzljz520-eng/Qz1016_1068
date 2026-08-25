package workflow

import (
	"fmt"
	"storeinspection/config"
	"storeinspection/model"
)

func ValidateRecord(r model.Record, p config.Policy) error {
	if r.ID == "" || r.StoreID == "" {
		return fmt.Errorf("record identity required")
	}
	if !p.AllowsStatus(r.Status) {
		return model.ErrInvalidStatus
	}
	if r.Severity < 1 || r.Severity > p.MaxSeverity {
		return fmt.Errorf("severity out of range")
	}
	if p.RequireAssignee && r.Assignee == "" {
		return fmt.Errorf("assignee required")
	}
	return nil
}
func NextStatus(current, action string) string {
	switch action {
	case "assign":
		return "assigned"
	case "start":
		return "in_progress"
	case "resolve":
		return "resolved"
	case "close":
		return "closed"
	case "archive":
		return "archived"
	}
	return current
}
