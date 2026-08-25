package service

import (
	"example.com/emblem/model"
	"example.com/emblem/storage"
	"path/filepath"
	"testing"
)

func TestServiceCreate(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	v := New(s)
	e := v.CreateTeam(model.Team{ID: "t", Slug: "t", Name: "T", Season: "2026", Slogan: "Go!", Active: true}, model.LogoAsset{ID: "t", TeamID: "t", ParticleCount: 20, Direction: 1}, model.Announcement{ID: "t", TeamID: "t", Title: "Hi", Body: "Go", Published: true})
	if e != nil {
		t.Fatal(e)
	}
}
