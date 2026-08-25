package animation

import "testing"

func TestEngineInputs(t *testing.T) {
	e := New(32)
	e.Apply(Space, 0, 0)
	if !e.Snapshot().Fractured {
		t.Fatal()
	}
	e.Apply(Left, 0, 0)
	if e.Snapshot().Direction != -1 {
		t.Fatal()
	}
}
