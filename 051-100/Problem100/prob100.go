package main

import (
	helper "MathLibGo"
	"fmt"
	"math"
)

const float64EqualityThreshold = 1e-9

var cached_IsSquare = helper.CacheWrapper_SingleToSingle(func(n float64) bool {
	sqrt := math.Sqrt(n)
	return math.Abs(sqrt-float64(int(sqrt))) < float64EqualityThreshold
})

func SolutionClosureB() func() uint64 {
	a := uint64(1)
	b := uint64(3)

	return func() uint64 {
		c := (6 * b) - a - 2
		a, b = b, c
		return c
	}
}

func SolutionClosureC() func() uint64 {
	a := uint64(0)
	b := uint64(1)

	return func() uint64 {
		c := (6 * b) - a
		a, b = b, c
		return c
	}
}

func main() {
	// The following gives the first n, b values which satisfy the arrangement
	// Taking the b values, we find OEIS sequence A011900 for the values of b, and A001109 for r
	// This leads to the subsequent solution using recurrence relations

	// n := 1.0
	// for true {
	// 	computed := float64(1 + (2 * n * (n - 1)))
	// 	b := (1.0 / 2) + (math.Sqrt(computed) / 2)
	// 	if cached_IsSquare(computed) && int(math.Sqrt(computed))%2 != 0 {
	// 		fmt.Printf("n = %f, b = %f\n", n, b)
	// 		fmt.Printf("(%d / %d) * (%d / %d) = %.20f\n", int(b), int(n), int(b-1), int(n-1), (b/n)*((b-1)/(n-1)))

	// 		if n > math.Pow10(12) {
	// 			break
	// 		}
	// 	}
	// 	n += 1
	// }

	b := SolutionClosureB()
	r := SolutionClosureC()

	b_val, r_val := b(), r()
	for true {
		b_val, r_val = b(), r()

		fmt.Println(b_val, r_val, b_val+r_val)

		if b_val+r_val > uint64(math.Pow10(12)) {
			break
		}
	}
}
