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
func (this Point[T]) move(d Direction[T]) Point[T] {
	return Point[T]{
		X: this.X + d.DX,
		Y: this.Y + d.DY,
	}
}

type Direction[T Numeric] struct {
	DX T
	DY T
}

func neighbors4[T Numeric]() []Direction[T] {
	return []Direction[T]{
		{DX: -1, DY: 0},
		{DX: 1, DY: 0},
		{DX: 0, DY: 1},
		{DX: 0, DY: -1},
	}
}
func CityBlockDistance[T Numeric](p, q Point[T]) int {
	x, y := Diff(p, q)
	return int(math.Abs(x)) + int(math.Abs(y))
}
func Diff[T Numeric](p, q Point[T]) (x, y float64) {
	return float64(p.X - q.X), float64(p.Y - q.Y)
}
