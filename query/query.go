package query

import (
	"example.com/emblem/animation"
	"example.com/emblem/model"
	"example.com/emblem/service"
)

type Launch struct {
	Team         model.Team
	Logo         model.LogoAsset
	Announcement model.Announcement
	Points       []animation.Point
}
type Query struct{ Service *service.Service }

func New(s *service.Service) *Query { return &Query{Service: s} }
func (q *Query) Launch(id string) (Launch, error) {
	t, e := q.Service.Team(id)
	if e != nil {
		return Launch{}, e
	}
	if !t.Active {
		return Launch{}, model.ErrInactive
	}
	l, e := q.Service.Store.GetLogo(id)
	if e != nil {
		return Launch{}, e
	}
	a, e := q.Service.Store.GetAnnouncement(id)
	if e != nil {
		return Launch{}, e
	}
	eng := animation.New(l.ParticleCount)
	return Launch{t, l, a, eng.Points()}, nil
}
func (q *Query) TeamSummary(id string) (string, error) { return q.Service.Summary(id) }
func (q *Query) Announcement(id string) (model.Announcement, error) {
	return q.Service.Store.GetAnnouncement(id)
}
