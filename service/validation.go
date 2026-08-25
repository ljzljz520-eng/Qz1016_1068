package service

import (
	"fmt"
	"storeinspection/config"
	"storeinspection/model"
)

func ValidateUser(u model.User) error {
	if u.ID == "" || u.Name == "" {
		return fmt.Errorf("user identity required")
	}
	switch u.Role {
	case "viewer", "manager", "auditor":
	default:
		return fmt.Errorf("unknown role")
	}
	return nil
}
func ValidatePolicy(p config.Policy) error {
	if p.MaxSeverity < 1 {
		return fmt.Errorf("max severity invalid")
	}
	if len(p.AllowedStatuses) == 0 {
		return fmt.Errorf("statuses required")
	}
	return nil
}
func ValidateTransition(r model.Record, next string) error {
	if next == r.Status {
		return fmt.Errorf("status unchanged")
	}
	if next == "archived" && !r.IsResolved() {
		return fmt.Errorf("record not resolved")
	}
	return nil
}
func IsStoreID(s string) bool  { return len(s) >= 2 && len(s) <= 32 }
func IsRecordID(s string) bool { return len(s) >= 3 && len(s) <= 64 }
func IsSeverity(n int) bool    { return n >= 1 && n <= 5 }
func IsStatus(s string) bool   { return config.DefaultPolicy().AllowsStatus(s) }
func NormalizeTitle(s string) string {
	if len(s) > 120 {
		return s[:120]
	}
	return s
}
