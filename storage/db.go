package storage

import (
	"encoding/json"
	"go.etcd.io/bbolt"
	"sync"
)

var recordsBucket = []byte("records")
var eventsBucket = []byte("events")
var auditsBucket = []byte("audits")
var usersBucket = []byte("users")

type DB struct {
	bolt *bbolt.DB
	mu   sync.RWMutex
}

func Open(path string) (*DB, error) {
	b, err := bbolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}
	d := &DB{bolt: b}
	err = d.init()
	if err != nil {
		b.Close()
		return nil, err
	}
	return d, nil
}
func (d *DB) init() error {
	return d.bolt.Update(func(tx *bbolt.Tx) error {
		for _, n := range [][]byte{recordsBucket, eventsBucket, auditsBucket, usersBucket} {
			if _, err := tx.CreateBucketIfNotExists(n); err != nil {
				return err
			}
		}
		return nil
	})
}
func (d *DB) Close() error         { d.mu.Lock(); defer d.mu.Unlock(); return d.bolt.Close() }
func encode(v any) ([]byte, error) { return json.Marshal(v) }
func decode(b []byte, v any) error { return json.Unmarshal(b, v) }
