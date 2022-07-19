package main

import "github.com/mlambda-net/hcl/lib"

func main() {

	items := lib.Sum([]int32{1, 2, 3, 4, 5, 6, 7, 8})
	println(items)

}
