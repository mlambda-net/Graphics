package abstract

type Algebra interface {
	Number[Numeric] | Symbolic[any]
}

type float interface {
	float32 | float64
}

type integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

type Numeric interface {
	integer | float
}
