package storage

import (
	"go.etcd.io/bbolt"
	"storeinspection/model"
	"time"
)

func (d *DB) Count(bucket []byte) (int, error) {
	n := 0
	d.mu.RLock()
	defer d.mu.RUnlock()
	e := d.bolt.View(func(t *bbolt.Tx) error {
		return t.Bucket(bucket).ForEach(func(k, v []byte) error {
			if v != nil {
				n++
			}
			return nil
		})
	})
	return n, e
}
func (d *DB) PurgeEvents(before time.Time) (int, error) {
	n := 0
	d.mu.RLock()
	defer d.mu.RUnlock()
	e := d.bolt.Update(func(t *bbolt.Tx) error {
		b := t.Bucket(eventsBucket)
		keys := [][]byte{}
		if err := b.ForEach(func(k, v []byte) error {
			var ev model.Event
			if decode(v, &ev) == nil && ev.At.Before(before) {
				keys = append(keys, append([]byte{}, k...))
			}
			return nil
		}); err != nil {
			return err
		}
		for _, k := range keys {
			if err := b.Delete(k); err != nil {
				return err
			}
			n++
		}
		return nil
	})
	return n, e
}
func (d *DB) HasRecord(id string) bool          { _, e := d.GetRecord(id); return e == nil }
func (d *DB) Snapshot() ([]model.Record, error) { return d.ListRecords("", "") }
