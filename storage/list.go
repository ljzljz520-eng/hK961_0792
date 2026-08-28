package storage

import (
	"example.com/emblem/model"
	"go.etcd.io/bbolt"
)

func (s *Store) ListTeams() ([]model.Team, error) {
	out := []model.Team{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("teams")).ForEach(func(_, v []byte) error {
			var t model.Team
			if e := model.Decode(v, &t); e != nil {
				return e
			}
			out = append(out, t)
			return nil
		})
	})
	return out, e
}
func (s *Store) ListEvents() ([]model.AuditEvent, error) {
	out := []model.AuditEvent{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("events")).ForEach(func(_, v []byte) error {
			var x model.AuditEvent
			if e := model.Decode(v, &x); e != nil {
				return e
			}
			out = append(out, x)
			return nil
		})
	})
	return out, e
}
func (s *Store) DeleteTeam(id string) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte("teams")).Delete([]byte(id)) })
}
