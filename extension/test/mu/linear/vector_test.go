package linear

import (
	"fmt"
	"github.com/mlambda-net/extension/mu/abstract"
	"testing"
)

type money struct {
	value float64
	name  string
}

func (m money) UnWrap() money {
	return m
}

func (m money) Times(other abstract.Symbolic[money]) abstract.Symbolic[money] {
	o := other.UnWrap()

	if m.name == other.UnWrap().name {
		return &money{
			value: m.value * o.value,
			name:  m.name,
		}
	}

	return &money{
		value: 0,
		name:  "",
	}

}

func (m money) Divide(other abstract.Symbolic[money]) abstract.Symbolic[money] {
	o := other.UnWrap()
	if m.name == other.UnWrap().name {
		return &money{
			value: m.value / o.value,
			name:  m.name,
		}
	}

	return &money{
		value: 0,
		name:  "",
	}
}

func (m money) Plus(other abstract.Symbolic[money]) abstract.Symbolic[money] {
	o := other.UnWrap()
	if m.name == other.UnWrap().name {
		return &money{
			value: m.value + o.value,
			name:  m.name,
		}
	}

	return &money{
		value: 0,
		name:  "",
	}
}

func (m money) Minus(other abstract.Symbolic[money]) abstract.Symbolic[money] {
	o := other
	if m.name == other.UnWrap().name {
		return &money{
			value: m.value - o.value,
			name:  m.name,
		}
	}

	return &money{
		value: 0,
		name:  "",
	}
}

func TestVectorFromNumeric(t *testing.T) {

	k := money{
		value: 100,
		name:  "USD",
	}
	r := money{
		value: 200,
		name:  "USD",
	}

	z := k.Plus(r).UnWrap()

	fmt.Printf("name: %s, value: %f", z.name, z.value)

}
