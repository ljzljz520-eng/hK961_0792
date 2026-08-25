package storage

import (
	"example.com/emblem/model"
	"go.etcd.io/bbolt"
	"path/filepath"
	"sync"
)

var buckets = []string{"teams", "logos", "announcements", "events"}

type Store struct {
	db *bbolt.DB
	mu sync.RWMutex
}

func Open(path string) (*Store, error) {
	db, e := bbolt.Open(filepath.Clean(path), 0600, nil)
	if e != nil {
		return nil, e
	}
	s := &Store{db: db}
	e = db.Update(func(tx *bbolt.Tx) error {
		for _, n := range buckets {
			if _, x := tx.CreateBucketIfNotExists([]byte(n)); x != nil {
				return x
			}
		}
		return nil
	})
	if e != nil {
		db.Close()
		return nil, e
	}
	return s, nil
}
func (s *Store) Close() error { s.mu.Lock(); defer s.mu.Unlock(); return s.db.Close() }
func (s *Store) put(bucket, key string, v any) error {
	b, e := model.Encode(v)
	if e != nil {
		return e
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte(bucket)).Put([]byte(key), b) })
}
func (s *Store) get(bucket, key string, v any) error {
	return s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket)).Get([]byte(key))
		if b == nil {
			return bbolt.ErrBucketNotFound
		}
		return model.Decode(append([]byte(nil), b...), v)
	})
}
func (s *Store) SaveTeam(v model.Team) error { return s.put("teams", v.ID, v) }
func (s *Store) GetTeam(id string) (model.Team, error) {
	var v model.Team
	e := s.get("teams", id, &v)
	return v, e
}
func (s *Store) SaveLogo(v model.LogoAsset) error { return s.put("logos", v.ID, v) }
func (s *Store) GetLogo(id string) (model.LogoAsset, error) {
	var v model.LogoAsset
	e := s.get("logos", id, &v)
	return v, e
}
func (s *Store) SaveAnnouncement(v model.Announcement) error { return s.put("announcements", v.ID, v) }
func (s *Store) GetAnnouncement(id string) (model.Announcement, error) {
	var v model.Announcement
	e := s.get("announcements", id, &v)
	return v, e
}
func (s *Store) SaveEvent(v model.AuditEvent) error { return s.put("events", v.ID, v) }
func (s *Store) GetEvent(id string) (model.AuditEvent, error) {
	var v model.AuditEvent
	e := s.get("events", id, &v)
	return v, e
}
