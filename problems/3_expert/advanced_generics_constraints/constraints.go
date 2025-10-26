package advconstraints

// Number is a type set of basic numeric types we support.
type Number interface { ~int | ~int64 | ~float64 | ~float32 }

func Add[T Number](a, b T) T { return a + b }
