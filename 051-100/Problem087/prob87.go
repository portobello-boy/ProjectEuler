package main

import (
	"MathLibGo/numberTheory/primes"
	"fmt"
	"math"
)

var target int = 50000000

func main() {
	prime_squares := make(map[int]int)
	prime_cubes := make(map[int]int)
	prime_fourths := make(map[int]int)

	primeList := primes.BoundedPrimeSequence(int(math.Ceil(math.Pow(float64(target), 1.0/2))))
	for _, p := range primeList {
		// fmt.Println(p)

		square := p * p
		cube := square * p
		fourth := cube * p

		prime_squares[p] = p * p
		if prime_squares[p] > target {
			break
		}
		if cube < target {
			prime_cubes[p] = cube
		}
		if fourth < target {
			prime_fourths[p] = fourth
		}
	}

	// Count all combinations of squares, cubes, and fourth powers with a sum less than the target
	summable_numbers := make(map[int]bool)
	for _, square := range prime_squares {
		for _, cube := range prime_cubes {
			for _, fourth := range prime_fourths {
				s := square + cube + fourth
				if s > target {
					continue
				}
				if _, ok := summable_numbers[s]; !ok {
					summable_numbers[s] = true
					// fmt.Printf("%d = %d + %d + %d and is a sum of squares, cubes, and fourth powers.\n", s, square, cube, fourth)

				}
			}
		}
	}
	fmt.Println(len(summable_numbers))
}
