package abstract



type symbolic[T Algebra] struct {
	item T
}

func (s symbolic[T]) Times(other symbolic[T]) Symbolic[T] {
	a := s.item + other.item

	s.item other.UnWrap()
	return symbolic[T]{item: other.UnWrap()}
}

func (s symbolic[T]) Divide(other Symbolic[T]) Symbolic[T] {
	//TODO implement me
	panic("implement me")
}

func (s symbolic[T]) Plus(other Symbolic[T]) Symbolic[T] {
	//TODO implement me
	panic("implement me")
}

func (s symbolic[T]) Minus(other Symbolic[T]) Symbolic[T] {
	//TODO implement me
	panic("implement me")
}

func (s symbolic[T]) UnWrap() T {
	return s.item
}

func ToSymbolic[T any](item T) Symbolic[T] {
	return symbolic[T]{item: item}
}
