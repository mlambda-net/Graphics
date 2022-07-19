package test

import (
	"fmt"
	"github.com/mlambda-net/hcl/lib"
	"testing"
)

func TestSum(t *testing.T) {

	var vector []int32

	for i := 0; i < 16384; i++ {
		vector = append(vector, int32(i))
	}

	x := lib.Sum(vector)
	println(x)
}

var table = []struct {
	input int
}{
	{input: 100},
	{input: 1000},
	{input: 10000},
	{input: 100000},
}

func BenchmarkParallel(b *testing.B) {

	var vector []int32
	for i := 0; i < 16384; i++ {
		vector = append(vector, int32(i))
	}
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				lib.Sum(vector)
			}
		})
	}
}

func BenchmarkSequential(b *testing.B) {

	var vector []int32
	for i := 0; i < 16384; i++ {
		vector = append(vector, int32(i))
	}

	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Sum(vector)
			}
		})
	}
}

func Sum(vector []int32) {
	var x int32 = 0
	for _, i := range vector {
		x += i
	}
}
