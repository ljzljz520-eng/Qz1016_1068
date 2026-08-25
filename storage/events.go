package storage

import (
	"go.etcd.io/bbolt"
	"sort"
	"storeinspection/model"
)

func (d *DB) PutEvent(e model.Event) error {
	b, err := encode(e)
	if err != nil {
		return err
	}
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.bolt.Update(func(tx *bbolt.Tx) error { return tx.Bucket(eventsBucket).Put([]byte(e.ID), b) })
}
func (d *DB) ListEvents(record string) ([]model.Event, error) {
	out := []model.Event{}
	d.mu.RLock()
	defer d.mu.RUnlock()
	err := d.bolt.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(eventsBucket).ForEach(func(_, v []byte) error {
			var e model.Event
			if err := decode(v, &e); err != nil {
				return err
			}
			if record == "" || e.RecordID == record {
				out = append(out, e)
			}
			return nil
		})
	})
	sort.Slice(out, func(i, j int) bool { return out[i].At.Before(out[j].At) })
	return out, err
}
func (d *DB) PutAudit(a model.Audit) error {
	b, err := encode(a)
	if err != nil {
		return err
	}
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.bolt.Update(func(tx *bbolt.Tx) error { return tx.Bucket(auditsBucket).Put([]byte(a.ID), b) })
}
func (d *DB) ListAudits(record string) ([]model.Audit, error) {
	out := []model.Audit{}
	d.mu.RLock()
	defer d.mu.RUnlock()
	err := d.bolt.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(auditsBucket).ForEach(func(_, v []byte) error {
			var a model.Audit
			if err := decode(v, &a); err != nil {
				return err
			}
			if record == "" || a.RecordID == record {
				out = append(out, a)
			}
			return nil
		})
	})
	return out, err
}
