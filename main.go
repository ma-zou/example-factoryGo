package main

import (
	"factoryPattern/shape"
	"fmt"
)

func main() {
	square := shape.ShapeFactory(func() shape.Square {
		return shape.NewSquare(10, 12)
	})

	triangle := shape.ShapeFactory(func() shape.Triangle {
		return shape.NewTriangle(10, 20)
	})

	circle := shape.ShapeFactory(func() shape.Circle {
		return shape.NewCircle(20)
	})

	fmt.Println(square.Draw(), "has a size of", square.GetSize())
	fmt.Println(triangle.Draw(), "has a size of", triangle.GetSize())
	fmt.Println(circle.Draw(), "has a size of", circle.GetSize())
}
