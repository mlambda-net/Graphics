package generic

type Pair[T any] interface {
	First() T
	Second() T
}

type pair[T any] struct {
	first  T
	second T
}

func (p *pair[T]) First() T {
	return p.first
}

func (p *pair[T]) Second() T {
	return p.second
}

func NewPair[T any](a T, b T) Pair[T] {
	return &pair[T]{first: a, second: b}
}
