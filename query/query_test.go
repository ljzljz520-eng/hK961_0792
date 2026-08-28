package query

import (
	"example.com/emblem/model"
	"example.com/emblem/service"
	"example.com/emblem/storage"
	"path/filepath"
	"testing"
)

func TestProjection(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "q"))
	defer s.Close()
	v := service.New(s)
	v.CreateTeam(model.Team{ID: "t", Slug: "t", Name: "T", Season: "26", Slogan: "Win", Active: true}, model.LogoAsset{ID: "t", TeamID: "t", ParticleCount: 20, Direction: 1}, model.Announcement{ID: "t", TeamID: "t", Title: "A", Body: "B", Published: true})
	if _, e := New(v).Launch("t"); e != nil {
		t.Fatal(e)
	}
}
