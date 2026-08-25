package service

import (
	"fmt"
	"storeinspection/model"
)

func Authorize(u model.User, action string) error {
	if !u.Active {
		return fmt.Errorf("inactive user")
	}
	if action == "review" && !u.CanReview() {
		return fmt.Errorf("review permission denied")
	}
	if action == "assign" && !u.CanAssign() {
		return fmt.Errorf("assign permission denied")
	}
	if action == "archive" && u.Role != "auditor" && u.Role != "manager" {
		return fmt.Errorf("archive permission denied")
	}
	return nil
}
func AllowedActions(u model.User) []string {
	if !u.Active {
		return nil
	}
	out := []string{"view"}
	if u.CanAssign() {
		out = append(out, "assign")
	}
	if u.CanReview() {
		out = append(out, "review", "resolve")
	}
	if u.Role == "manager" {
		out = append(out, "archive")
	}
	return out
}
