package main

import "fmt"

/*
   double pop_mean(int numPoints, double a[]) {
       if (a == NULL || numPoints == 0) {
           return 0;
       }
       double mean = 0;
       for (int i = 0; i < numPoints; i++) {
           mean+=a[i];
       }
       return mean / numPoints;
   }
*/
import "C"

func popMean(a []float64) float64 {
	// This is the general case, which includes the special cases
	// of zero-value (a == nil and len(a) == 0)
	// and zero-length (len(a) == 0) slices.
	if len(a) == 0 {
		return 0
	}

	return float64(C.pop_mean(C.int(len(a)), (*C.double)(&a[0])))
}

func main() {
	a := make([]float64, 10)
	for i := range a {
		a[i] = float64(i + 1)
	}

	// slice
	fmt.Println(len(a), a)
	pm := popMean(a)
	fmt.Println(pm)

	// subslice
	b := a[1:4]
	fmt.Println(len(b), b)
	pm = popMean(b)
	fmt.Println(pm)

	// zero length
	c := a[:0]
	fmt.Println(len(c), c)
	pm = popMean(c)
	fmt.Println(pm)

	// zero value (nil)
	var z []float64
	fmt.Println(len(z), z, z == nil)
	pm = popMean(z)
	fmt.Println(pm)
}
