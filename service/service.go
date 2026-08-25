package service

import (
	"example.com/emblem/cache"
	"example.com/emblem/model"
	"example.com/emblem/storage"
	"fmt"
)

type Service struct {
	Store *storage.Store
	Cache *cache.Cache
}

func New(s *storage.Store) *Service { return &Service{Store: s, Cache: cache.New()} }
func (s *Service) CreateTeam(t model.Team, l model.LogoAsset, a model.Announcement) error {
	if e := t.Validate(); e != nil {
		return e
	}
	if l.TeamID != t.ID {
		l.TeamID = t.ID
	}
	if e := l.Validate(); e != nil {
		return e
	}
	if a.TeamID != t.ID {
		a.TeamID = t.ID
	}
	if e := a.Validate(); e != nil {
		return e
	}
	if e := s.Store.SaveTeam(t); e != nil {
		return e
	}
	if e := s.Store.SaveLogo(l); e != nil {
		return e
	}
	return s.Store.SaveAnnouncement(a)
}
func (s *Service) UpdateTeam(t model.Team) error {
	if e := t.Validate(); e != nil {
		return e
	}
	if e := s.Store.SaveTeam(t); e != nil {
		return e
	}
	return nil
}
func (s *Service) SaveAnnouncement(a model.Announcement) error {
	if e := a.Validate(); e != nil {
		return e
	}
	return s.Store.SaveAnnouncement(a)
}
func (s *Service) RecordEvent(e model.AuditEvent) error {
	if e.CreatedAt == "" {
		e.CreatedAt = "deterministic"
	}
	return s.Store.SaveEvent(e)
}
func (s *Service) Team(id string) (model.Team, error) {
	if v, ok := s.Cache.Get("team:" + id); ok {
		return v.(model.Team), nil
	}
	v, e := s.Store.GetTeam(id)
	if e == nil {
		s.Cache.Put("team:"+id, v)
	}
	return v, e
}
func (s *Service) RefreshTeam(id string) (model.Team, error) {
	s.Cache.Invalidate("team:" + id)
	return s.Team(id)
}
func (s *Service) ValidateAndUpdate(t model.Team) error { return s.UpdateTeam(t) }
func (s *Service) Summary(id string) (string, error) {
	t, e := s.Team(id)
	if e != nil {
		return "", e
	}
	return fmt.Sprintf("%s | %s | %s", t.Name, t.Season, t.Slogan), nil
}
