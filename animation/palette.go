package animation

type Palette struct{ Primary, Secondary, Accent string }

func DefaultPalette() Palette { return Palette{"#ef3340", "#111827", "#f8fafc"} }
func Blend(a, b Palette, ratio float64) Palette {
	if ratio < 0 {
		ratio = 0
	}
	if ratio > 1 {
		ratio = 1
	}
	if ratio < 0.5 {
		return a
	}
	return b
}
func (p Palette) Valid() bool { return p.Primary != "" && p.Secondary != "" && p.Accent != "" }
