package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

var ErrInactive = errors.New("team inactive")

type Team struct {
	ID, Slug, Name, Season, Slogan, Color string
	Active                                bool
}
type LogoAsset struct {
	ID, TeamID, Shape string
	ParticleCount     int
	RotationSpeed     float64
	Direction         int
}
type Announcement struct {
	ID, TeamID, Title, Body string
	Published               bool
}
type AuditEvent struct{ ID, TeamID, Event, Payload, CreatedAt string }

func (t Team) Validate() error {
	if strings.TrimSpace(t.ID) == "" || strings.TrimSpace(t.Name) == "" {
		return errors.New("team id and name required")
	}
	if strings.ContainsAny(t.Slug, " !@#$%^&*()") {
		return errors.New("invalid slug")
	}
	if len(t.Slogan) < 3 {
		return errors.New("slogan too short")
	}
	if t.Season == "" {
		return errors.New("season required")
	}
	return nil
}
func (l LogoAsset) Validate() error {
	if l.ID == "" || l.TeamID == "" {
		return errors.New("logo identity required")
	}
	if l.ParticleCount < 16 || l.ParticleCount > 10000 {
		return errors.New("particle count out of range")
	}
	if l.Direction != 1 && l.Direction != -1 {
		return errors.New("direction must be one or minus one")
	}
	return nil
}
func (a Announcement) Validate() error {
	if a.ID == "" || a.TeamID == "" || a.Title == "" || a.Body == "" {
		return errors.New("announcement fields required")
	}
	return nil
}
func (e AuditEvent) Validate() error {
	if e.ID == "" || e.TeamID == "" || e.Event == "" {
		return errors.New("event fields required")
	}
	return nil
}
func Encode(v any) ([]byte, error)     { return json.Marshal(v) }
func Decode(b []byte, v any) error     { return json.Unmarshal(b, v) }
func TeamKey(id string) string         { return fmt.Sprintf("team:%s", id) }
func LogoKey(id string) string         { return fmt.Sprintf("logo:%s", id) }
func AnnouncementKey(id string) string { return fmt.Sprintf("announcement:%s", id) }
func EventKey(id string) string        { return fmt.Sprintf("event:%s", id) }
