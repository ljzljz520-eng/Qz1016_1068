package audit

import (
	"fmt"
	"storeinspection/model"
	"storeinspection/storage"
)

type Logger struct{ db *storage.DB }

func New(db *storage.DB) *Logger { return &Logger{db: db} }
func (l *Logger) Record(action, record, actor string) error {
	a := model.NewAudit(fmt.Sprintf("audit-%d", len(action)+len(record)+len(actor)), action, record, actor)
	return l.db.PutAudit(a)
}
func (l *Logger) History(record string) ([]model.Audit, error) { return l.db.ListAudits(record) }
