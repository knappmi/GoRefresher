package shape

// Shape defines area & perimeter behavior.
type Shape interface { Area() float64; Perimeter() float64 }

type Rectangle struct { W, H float64 }
func (r Rectangle) Area() float64 { // TODO implement
	return 0 }
func (r Rectangle) Perimeter() float64 { // TODO implement
	return 0 }

type Circle struct { R float64 }
func (c Circle) Area() float64 { // TODO implement
	return 0 }
func (c Circle) Perimeter() float64 { // TODO implement
	return 0 }

type Triangle struct { A,B,C float64 }
func (t Triangle) Perimeter() float64 { // TODO implement
	return 0 }
func (t Triangle) Area() float64 { // TODO implement (Heron's formula)
	return 0 }

func SumAreas(ss []Shape) float64 { // TODO implement aggregation
	return 0 }
