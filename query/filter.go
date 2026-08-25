package query

import (
	"example.com/emblem/model"
	"sort"
)

func ActiveTeams(in []model.Team) []model.Team {
	out := []model.Team{}
	for _, t := range in {
		if t.Active {
			out = append(out, t)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Slug < out[j].Slug })
	return out
}
func FindBySlug(in []model.Team, slug string) (model.Team, bool) {
	for _, t := range in {
		if t.Slug == slug {
			return t, true
		}
	}
	return model.Team{}, false
}
func PublishedAnnouncements(in []model.Announcement) []model.Announcement {
	out := []model.Announcement{}
	for _, a := range in {
		if model.IsPublished(a) {
			out = append(out, a)
		}
	}
	return out
}
