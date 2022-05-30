package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

const pi = math.Pi

func (circle Circle) CalcArea() float64 {
	return math.Pow(circle.Radius, 2) * pi
}

func (circle Circle) CalcPerimeter() float64 {
	return 2 * pi * circle.Radius
}
