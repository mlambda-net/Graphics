package queryable

import (
	"github.com/mlambda-net/extension/mu/generic"
)

type Option[T any] interface {
	Subscribe(func(T))
	UnWrap() (T, error)
}

type optionError[T any] struct {
	item T
}

type optionJust[T any] struct {
	item T
}

func Cast[T any, Y any](x T) Y {

	y, ok := any(x).(Y)
	if ok {
		return y
	}

	return y
}

type IEnumerable[T any] interface {
	ForEach(fn func(c T))
	Map(fn func(c T) T) IEnumerable[T]
	ToArray() []T
	Aggregate(fn func(c, a T) T) T
	Count() int
	Zip(items IEnumerable[T]) IEnumerable[generic.Pair[any]]
	Get(index int) T
	FlatMap(func(T, any) any)
}

type enumerable[T any] struct {
	items []T
}

func (i *enumerable[T]) Count() int {
	return len(i.items)
}

func (i *enumerable[T]) Aggregate(fn func(c T, a T) T) T {
	var accumulator T
	i.ForEach(func(c T) {
		accumulator = fn(c, accumulator)
	})
	return accumulator
}

func (i *enumerable[T]) ToArray() []T {
	return i.items
}

func (i *enumerable[T]) Get(index int) T {

	var def T
	if index <= len(i.items) {
		return i.items[index]
	}

	return def
}

func (i *enumerable[T]) Zip(items IEnumerable[T]) IEnumerable[generic.Pair[any]] {
	if i.Count() > items.Count() {
		return i.zip(items)
	}
	return items.Zip(i)
}

func (i *enumerable[T]) zip(items IEnumerable[T]) IEnumerable[generic.Pair[any]] {
	var pairs []generic.Pair[any]
	for k, i := range i.items {
		pairs = append(pairs, generic.NewPair[any](i, items.Get(k)))
	}
	return Enumerable(pairs)
}

func (i *enumerable[T]) Map(fn func(c T) T) IEnumerable[T] {
	var items []T
	i.ForEach(func(c T) {
		items = append(items, fn(c))
	})

	return &enumerable[T]{items: items}
}

func (i *enumerable[T]) ForEach(fn func(item T)) {
	for _, i := range i.items {
		fn(i)
	}
}

func Enumerable[T any](items []T) IEnumerable[T] {
	return &enumerable[T]{
		items: items,
	}
}
