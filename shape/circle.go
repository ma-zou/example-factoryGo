package shape

import (
	"math"
)

type Circle struct {
	radius float64
}

func NewCircle(radius float64) Circle {
	return Circle{radius}
}

func (c Circle) Draw() string {
	return "•"
}

func (c Circle) GetSize() float64 {
	return (c.radius * 2) * math.Pi
}
