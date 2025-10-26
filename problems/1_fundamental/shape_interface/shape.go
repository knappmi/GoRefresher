package shape

import "math"

// Shape defines area & perimeter behavior.
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	W, H float64
}

func (r Rectangle) Area() float64 {
	return r.W * r.H
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.W + r.H)
}

type Circle struct {
	R float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.R * c.R
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}

type Triangle struct {
	A, B, C float64
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

func (t Triangle) Area() float64 {
	s := t.Perimeter() / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func SumAreas(ss []Shape) float64 {
	var total float64
	for _, s := range ss {
		total += s.Area()
	}
	return total
}
