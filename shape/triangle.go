package shape

type Triangle struct {
	base   float64
	height float64
}

func NewTriangle(base, height float64) Triangle {
	return Triangle{base, height}
}

func (t Triangle) Draw() string {
	return "▲"
}

func (t Triangle) GetSize() float64 {
	return (t.base * t.height) / 2
}
