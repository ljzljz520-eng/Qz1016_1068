package service

import (
	"context"
	"fmt"
	"storeinspection/model"
)

type Notification struct {
	RecordID  string
	Recipient string
	Message   string
}

func (s *Service) Notify(ctx context.Context, id, recipient, message string) (Notification, error) {
	if err := ctx.Err(); err != nil {
		return Notification{}, err
	}
	if recipient == "" || message == "" {
		return Notification{}, fmt.Errorf("notification fields required")
	}
	if _, e := s.DB.GetRecord(id); e != nil {
		return Notification{}, e
	}
	return Notification{id, recipient, message}, nil
}
func FormatStatus(r model.Record) string {
	return fmt.Sprintf("%s is %s (severity %d)", r.Title, r.Status, r.Severity)
}
