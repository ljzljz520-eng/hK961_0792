package renderer

import "example.com/emblem/animation"

type Frame struct {
	Width, Height int
	Points        []animation.Point
	Opacity       float64
	Label         string
}
type Renderer struct{ width, height int }

func New(w, h int) *Renderer {
	if w < 1 {
		w = 1
	}
	if h < 1 {
		h = 1
	}
	return &Renderer{w, h}
}
func (r *Renderer) Dimensions() (int, int) { return r.width, r.height }
func (r *Renderer) Render(points []animation.Point, label string, opacity float64) Frame {
	if opacity < 0 {
		opacity = 0
	}
	if opacity > 1 {
		opacity = 1
	}
	out := append([]animation.Point(nil), points...)
	return Frame{r.width, r.height, out, opacity, label}
}
func (r *Renderer) Scale(p animation.Point) animation.Point {
	return animation.Point{p.X * float64(r.width) / 2, p.Y * float64(r.height) / 2, p.Z}
}
func (r *Renderer) Project(points []animation.Point) []animation.Point {
	out := make([]animation.Point, len(points))
	for i, p := range points {
		out[i] = r.Scale(p)
	}
	return out
}
