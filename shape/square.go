package shape

type Square struct {
	sideX float64
	sideY float64
}

func NewSquare(sideX, sideY float64) Square {
	return Square{sideX, sideY}
}

func (s Square) Draw() string {
	return "◼"
}

func (s Square) GetSize() float64 {
	return s.sideX * s.sideY
}
