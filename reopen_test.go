package arena

import (
	"example.com/emblem/model"
	"example.com/emblem/storage"
	"path/filepath"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "persist.db")
	s, e := storage.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	if e = s.SaveTeam(model.Team{ID: "persist", Slug: "persist", Name: "Persist", Season: "26", Slogan: "Keep", Active: true}); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = storage.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	if _, e = s.GetTeam("persist"); e != nil {
		t.Fatal(e)
	}
}
