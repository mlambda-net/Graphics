package abstract

type identity[T any] interface {
	UnWrap() T
}

type group[T any] interface {
	Times(other T) T
	Divide(other T) T
}

type ring[T any] interface {
	Plus(other T) T
	Minus(other T) T
}

type field[T any] interface {
	group[T]
	ring[T]
}
