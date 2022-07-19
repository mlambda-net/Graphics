package generic

import (
	"github.com/mlambda-net/extension/mu/queryable"
	"testing"
)
import "github.com/stretchr/testify/assert"

func Test_Map(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7}
	i := queryable.Enumerable[int](items)

	result := i.Map(func(c int) int {
		return c * 2
	})

	expected := []int{2, 4, 6, 8, 10, 12, 14}
	assert.Equal(t, expected, result.ToArray())
}

func Test_ForEach(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	i := queryable.Enumerable[int](items)

	actual := 0
	i.ForEach(func(c int) {
		actual += c
	})

	expected := 15
	assert.Equal(t, expected, actual)
}

func Test_Aggregate(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	i := queryable.Enumerable[int](items)

	actual := i.Aggregate(func(c int, accumulate int) int {
		return accumulate + c
	})

	expected := 15
	assert.Equal(t, expected, actual)
}
