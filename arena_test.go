package arena

import (
	"example.com/emblem/model"
	"example.com/emblem/query"
	"example.com/emblem/service"
	"example.com/emblem/storage"
	"path/filepath"
	"testing"
)

func fixture(t *testing.T) (*service.Service, *query.Query) {
	s, e := storage.Open(filepath.Join(t.TempDir(), "arena.db"))
	if e != nil {
		t.Fatal(e)
	}
	v := service.New(s)
	if e = v.CreateTeam(model.Team{ID: "storm", Slug: "storm", Name: "Storm", Season: "2026", Slogan: "Rise together", Active: true}, model.LogoAsset{ID: "storm", TeamID: "storm", Shape: "crest", ParticleCount: 64, Direction: 1}, model.Announcement{ID: "storm", TeamID: "storm", Title: "Reveal", Body: "We rise", Published: true}); e != nil {
		t.Fatal(e)
	}
	t.Cleanup(func() { s.Close() })
	return v, query.New(v)
}
func TestWorkflowLaunch(t *testing.T) {
	_, q := fixture(t)
	p, e := q.Launch("storm")
	if e != nil || len(p.Points) != 64 || p.Team.Slogan != "Rise together" {
		t.Fatal(e)
	}
}
func TestWorkflowManagement(t *testing.T) {
	v, q := fixture(t)
	if e := v.UpdateSlogan("storm", "New season"); e != nil {
		t.Fatal(e)
	}
	if _, e := q.TeamSummary("storm"); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowCustomize(t *testing.T) {
	v, _ := fixture(t)
	e := v
	_ = e
	if v.Cache.Size() != 0 {
		t.Fatal()
	}
}
func TestBusinessChain31(t *testing.T) {
	v, q := fixture(t)
	old, _ := q.TeamSummary("storm")
	if e := v.UpdateSlogan("storm", "Latest slogan"); e != nil {
		t.Fatal(e)
	}
	latest, _ := q.TeamSummary("storm")
	if latest == old {
		t.Fatalf("stale cached content: %q", latest)
	}
}
