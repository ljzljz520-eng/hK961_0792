package service

import (
	"example.com/emblem/model"
	"fmt"
)

func (s *Service) LaunchAudit(teamID string) error {
	return s.RecordEvent(model.AuditEvent{ID: "launch-" + teamID, TeamID: teamID, Event: "launch", Payload: "reveal"})
}
func (s *Service) UpdateSlogan(id, slogan string) error {
	t, e := s.Store.GetTeam(id)
	if e != nil {
		return e
	}
	t.Slogan = slogan
	return s.UpdateTeam(t)
}
func (s *Service) Publish(id string, published bool) error {
	a, e := s.Store.GetAnnouncement(id)
	if e != nil {
		return e
	}
	a.Published = published
	return s.SaveAnnouncement(a)
}
func (s *Service) Describe(id string) (string, error) {
	t, e := s.Team(id)
	if e != nil {
		return "", e
	}
	return fmt.Sprintf("%s/%s", t.Slug, t.Season), nil
}
func (s *Service) Invalidate(id string) { s.Cache.Invalidate("team:" + id) }
