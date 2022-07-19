package linear

import (
	"github.com/mlambda-net/extension/mu/abstract"
	"github.com/mlambda-net/extension/mu/queryable"
)

type Vector[T abstract.Numeric] interface {
	Sum() abstract.Number[T]
	Avg() abstract.Number[T]
	Count() abstract.Number[T]
	MultiplyByScalar(scalar abstract.Number[T])
	Multiply(vector Vector[T]) Vector[T]
	Product(vector Vector[T]) Matrix[T]
}

type vector[T abstract.Numeric] struct {
	items queryable.IEnumerable[abstract.Number[T]]
}

func (v *vector[T]) Count() abstract.Number[T] {
	count := v.items.Count()
	return abstract.ToNumber[T](T(count))
}

func (v *vector[T]) Sum() abstract.Number[T] {
	return v.items.Aggregate(func(c abstract.Number[T], a abstract.Number[T]) abstract.Number[T] {
		return a.Plus(c)
	})
}

func (v *vector[T]) Avg() abstract.Number[T] {
	avg := v.Sum().Divide(v.Count())
	return avg
}

func (v *vector[T]) Multiply(vector vector[T]) Vector[T] {
	r := v.items.Zip(vector.items)
	r.Select()
	return nil
}

func (v *vector[T]) MultiplyByScalar(scalar abstract.Number[T]) vector[T] {
	v.items.Map(func(c abstract.Number[T]) abstract.Number[T] {
		return scalar.Times(c)
	})
}

func (v *vector[T]) Product(vector Vector[T]) {
	//TODO implement me
	panic("implement me")
}

func ToNumericVector[T abstract.Numeric](items []T) Vector[T] {
	m := abstract.ToNumberEnumerable[T](items)
	return &vector[T]{items: m}
}
