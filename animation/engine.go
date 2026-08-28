package animation

import "math"

type Point struct{ X, Y, Z float64 }
type Input int

const (
	Drag Input = iota
	Space
	Left
	Right
	Up
	Down
)

type State struct {
	Yaw, Pitch float64
	Direction  int
	Fractured  bool
	Frame      uint64
}
type Engine struct {
	state  State
	points []Point
}

func New(count int) *Engine {
	if count < 1 {
		count = 1
	}
	p := make([]Point, count)
	for i := range p {
		a := float64(i) * 0.37
		p[i] = Point{math.Cos(a), math.Sin(a), math.Sin(a * 0.5)}
	}
	return &Engine{points: p, state: State{Direction: 1}}
}
func (e *Engine) Apply(in Input, dx, dy float64) {
	switch in {
	case Drag:
		e.state.Yaw += dx
		e.state.Pitch += dy
	case Space:
		e.state.Fractured = !e.state.Fractured
	case Left:
		e.state.Direction = -1
	case Right:
		e.state.Direction = 1
	case Up:
		e.state.Pitch -= 0.1
	case Down:
		e.state.Pitch += 0.1
	}
	e.state.Frame++
}
func (e *Engine) Snapshot() State { return e.state }
func (e *Engine) Points() []Point {
	out := make([]Point, len(e.points))
	copy(out, e.points)
	return out
}
func (e *Engine) Tick() { e.state.Yaw += 0.02 * float64(e.state.Direction); e.state.Frame++ }
func NormalizeAngle(v float64) float64 {
	for v > math.Pi {
		v -= 2 * math.Pi
	}
	for v < -math.Pi {
		v += 2 * math.Pi
	}
	return v
}
