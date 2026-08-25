package workflow

import (
	"fmt"
	"storeinspection/config"
	"storeinspection/model"
)

type Gate struct{ Policy config.Policy }

func NewGate(p config.Policy) Gate        { return Gate{Policy: p} }
func (g Gate) Check(r model.Record) error { return ValidateRecord(r, g.Policy) }
func (g Gate) CanTransition(r model.Record, next string) bool {
	if !g.Policy.AllowsStatus(next) {
		return false
	}
	if r.Status == "archived" {
		return false
	}
	if next == "archived" && (r.Status != "closed" && r.Status != "resolved") {
		return false
	}
	return true
}
func (g Gate) Transition(r *model.Record, next string) error {
	if !g.CanTransition(*r, next) {
		return fmt.Errorf("transition denied")
	}
	r.Status = next
	r.Version++
	return nil
}
func (g Gate) RequiredAssignee(next string) bool {
	return g.Policy.RequireAssignee && (next == "assigned" || next == "in_progress")
}
func (g Gate) ValidateActor(actor string) error {
	if actor == "" {
		return fmt.Errorf("actor required")
	}
	return nil
}
