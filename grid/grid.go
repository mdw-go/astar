package grid

import (
	"fmt"
	"math"
)

type Numeric interface{ int | float64 }

type Point[T Numeric] struct {
	X T
	Y T
}

func NewPoint[T Numeric](x, y T) Point[T] {
	return Point[T]{
		X: x,
		Y: y,
	}
}
func (this Point[T]) String() string {
	return fmt.Sprintf("[%v,%v]", this.X, this.Y)
}
func (this Point[T]) move(d direction[T]) Point[T] {
	return Point[T]{
		X: this.X + d.dX,
		Y: this.Y + d.dY,
	}
}

type direction[T Numeric] struct {
	dX T
	dY T
}

func neighbors4[T Numeric]() []direction[T] {
	return []direction[T]{
		{dX: -1, dY: 0},
		{dX: 1, dY: 0},
		{dX: 0, dY: 1},
		{dX: 0, dY: -1},
	}
}
func cityBlockDistance[T Numeric](p, q Point[T]) int {
	x, y := diff(p, q)
	return int(math.Abs(x)) + int(math.Abs(y))
}
func diff[T Numeric](p, q Point[T]) (x, y float64) {
	return float64(p.X - q.X), float64(p.Y - q.Y)
}
