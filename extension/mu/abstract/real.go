package abstract

import "github.com/mlambda-net/extension/mu/queryable"

type realNumber[T Numeric] struct {
	item T
}

func (r realNumber[T]) Times(other Number[T]) Number[T] {
	return &realNumber[T]{item: r.item * other.UnWrap()}
}

func (r realNumber[T]) Divide(other Number[T]) Number[T] {
	return &realNumber[T]{item: r.item / other.UnWrap()}
}

func (r realNumber[T]) Plus(other Number[T]) Number[T] {
	return &realNumber[T]{item: r.item / other.UnWrap()}
}

func (r realNumber[T]) Minus(other Number[T]) Number[T] {
	return &realNumber[T]{item: r.item / other.UnWrap()}
}

func (r realNumber[T]) UnWrap() T {
	return r.item
}

func ToNumber[T Numeric](item T) Number[T] {
	return realNumber[T]{item: item}
}

func ToNumberEnumerable[T Numeric](items []T) queryable.IEnumerable[Number[T]] {
	var array []Number[T]
	for _, i := range items {
		array = append(array, ToNumber(i))
	}
	return queryable.Enumerable(array)
}
