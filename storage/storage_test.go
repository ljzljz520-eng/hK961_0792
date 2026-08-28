package storage

import (
	"example.com/emblem/model"
	"path/filepath"
	"testing"
)

func TestStoreRoundTrip(t *testing.T) {
	p := filepath.Join(t.TempDir(), "a.db")
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	x := model.Team{ID: "t", Slug: "t", Name: "Team", Season: "S", Slogan: "Win!", Active: true}
	if e = s.SaveTeam(x); e != nil {
		t.Fatal(e)
	}
	if y, e := s.GetTeam("t"); e != nil || y.Name != x.Name {
		t.Fatal(y, e)
	}
}
