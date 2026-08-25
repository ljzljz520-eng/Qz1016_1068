package workflow

import (
	"context"
	"fmt"
	"storeinspection/model"
)

func CompleteChecklist(ctx context.Context, c model.Checklist) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if !c.RequiredComplete() {
		return fmt.Errorf("required checklist items incomplete")
	}
	return nil
}
func ChecklistProgress(c model.Checklist) float64 {
	if len(c.Items) == 0 {
		return 1
	}
	return float64(c.CompleteCount()) / float64(len(c.Items))
}
func ChecklistWarnings(c model.Checklist) []string {
	out := []string{}
	for _, i := range c.Items {
		if i.Required && !i.Complete {
			out = append(out, i.Label)
		}
	}
	return out
}
func ValidateChecklist(c model.Checklist) error {
	if c.RecordID == "" {
		return fmt.Errorf("record id required")
	}
	seen := map[string]bool{}
	for _, i := range c.Items {
		if i.Code == "" || seen[i.Code] {
			return fmt.Errorf("invalid checklist code")
		}
		seen[i.Code] = true
	}
	return nil
}
