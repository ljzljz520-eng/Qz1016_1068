package storage

import (
	"go.etcd.io/bbolt"
	"storeinspection/model"
)

func (d *DB) PutRecord(r model.Record) error {
	b, err := encode(r)
	if err != nil {
		return err
	}
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.bolt.Update(func(tx *bbolt.Tx) error { return tx.Bucket(recordsBucket).Put([]byte(r.ID), b) })
}
func (d *DB) GetRecord(id string) (model.Record, error) {
	var r model.Record
	d.mu.RLock()
	defer d.mu.RUnlock()
	err := d.bolt.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket(recordsBucket).Get([]byte(id))
		if v == nil {
			return model.ErrNotFound
		}
		return decode(v, &r)
	})
	return r, err
}
func (d *DB) ListRecords(store, status string) ([]model.Record, error) {
	out := []model.Record{}
	d.mu.RLock()
	defer d.mu.RUnlock()
	err := d.bolt.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(recordsBucket).ForEach(func(_, v []byte) error {
			if v == nil {
				return nil
			}
			var r model.Record
			if err := decode(v, &r); err != nil {
				return err
			}
			if (store == "" || r.StoreID == store) && (status == "" || r.Status == status) {
				out = append(out, r)
			}
			return nil
		})
	})
	return out, err
}
func (d *DB) DeleteRecord(id string) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.bolt.Update(func(tx *bbolt.Tx) error { return tx.Bucket(recordsBucket).Delete([]byte(id)) })
}
