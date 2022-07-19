package abstract

type Number[T Numeric] interface {
	field[Number[T]]
	identity[T]
}

type Symbolic[T any] interface {
	field[Symbolic[T]]
	identity[T]
}
