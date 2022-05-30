package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func (rectangle Rectangle) CalcArea() float64 {
	return rectangle.Weight * rectangle.Height
}

func (rectangle Rectangle) CalcPerimeter() float64 {
	return (rectangle.Weight + rectangle.Height) * 2
}
