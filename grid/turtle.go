package grid

import "github.com/mdwhatcott/astar"

type turtle[T Numeric] struct {
	grid map[Point[T]]struct{}
	from Point[T]
	to   Point[T]
}

func NewGridTurtle[T Numeric](grid map[Point[T]]struct{}, from, to Point[T]) astar.Turtle {
	return &turtle[T]{
		grid: grid,
		from: from,
		to:   to,
	}
}
func (this *turtle[T]) Search() (path []astar.Turtle, found bool) {
	return astar.SearchFrom(this)
}
func (this *turtle[T]) EstimatedDistanceToTarget() float64 {
	return float64(CityBlockDistance(this.from, this.to))
}
func (this *turtle[T]) AdjacentPositions() (results []astar.Turtle) {
	for _, d := range neighbors4[T]() {
		at := this.from.move(d)
		if _, contains := this.grid[at]; contains {
			results = append(results, NewGridTurtle(this.grid, at, this.to))
		}
	}
	return results
}
func (this *turtle[T]) StepCost() float64 { return 1 }
func (this *turtle[T]) Hash() string      { return this.from.String() }
func (this *turtle[T]) At() Point[T]      { return this.from }
