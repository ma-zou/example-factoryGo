package shape

type Shape interface {
	GetSize() float64
	Draw() string
}

func ShapeFactory[T Shape](newShape func() T) T {
	return newShape()
}
