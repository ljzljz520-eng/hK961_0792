package model

import "testing"

func TestTeamValidation(t *testing.T) {
	if (Team{ID: "x", Name: "N", Slug: "x", Season: "S", Slogan: "win"}).Validate() != nil {
		t.Fatal("valid team rejected")
	}
	if (Team{ID: "x"}).Validate() == nil {
		t.Fatal("invalid team accepted")
	}
}
