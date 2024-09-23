package shapes

import "math"

// Shape is implemented by anything that can tell us its area
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns the area of a rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter returns the perimeter of a rectangle.
func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

// Area returns the area of a circle.
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}
