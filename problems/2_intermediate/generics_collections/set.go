package genericscollections

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] { // TODO: implement set initialization
	return nil
}
func (s Set[T]) Add(v T) { // TODO: implement add
}
func (s Set[T]) Has(v T) bool { // TODO: implement membership check
	return false
}
