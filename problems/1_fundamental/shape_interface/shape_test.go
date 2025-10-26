package shape

import "testing"
import "math"

func TestAreas(t *testing.T) {
	if (Rectangle{2,3}).Area()!=6 { t.Fatal("rectangle area") }
	if math.Abs((Circle{1}).Area() - math.Pi) > 1e-9 { t.Fatal("circle area") }
	tri := Triangle{3,4,5}
	if math.Abs(tri.Area()-6) > 1e-9 { t.Fatalf("triangle area %f", tri.Area()) }
	sum := SumAreas([]Shape{Rectangle{2,3}, Circle{1}})
	if math.Abs(sum-(6+math.Pi))>1e-9 { t.Fatal("sum areas") }
}
