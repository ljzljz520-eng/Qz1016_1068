package storage

import (
	"go.etcd.io/bbolt"
	"storeinspection/model"
)

func (d *DB) PutUser(u model.User) error {
	b, e := encode(u)
	if e != nil {
		return e
	}
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.bolt.Update(func(t *bbolt.Tx) error { return t.Bucket(usersBucket).Put([]byte(u.ID), b) })
}
func (d *DB) GetUser(id string) (model.User, error) {
	var u model.User
	d.mu.RLock()
	defer d.mu.RUnlock()
	e := d.bolt.View(func(t *bbolt.Tx) error {
		v := t.Bucket(usersBucket).Get([]byte(id))
		if v == nil {
			return model.ErrNotFound
		}
		return decode(v, &u)
	})
	return u, e
}
func (d *DB) ListUsers() (out []model.User, e error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	e = d.bolt.View(func(t *bbolt.Tx) error {
		return t.Bucket(usersBucket).ForEach(func(_, v []byte) error {
			var u model.User
			if x := decode(v, &u); x != nil {
				return x
			}
			out = append(out, u)
			return nil
		})
	})
	return
}
