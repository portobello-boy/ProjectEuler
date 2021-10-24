package main

import (
	"os"
	"fmt"
	"strconv"
)

// I knew there was a way to generate pythagorean triples, so this function finds
// a triple (m, n, k) such that the generation of a, b, and c using that triple
// results in a*b*c = num. Because m and n must be coprime and with different parity,
// then this generation is much faster.
func approach1(num int) int {
	for m := 2; m < num; m++ {
		for n := m - 1; n > 0; n -= 2 {
			// Ensure m, n are coprime
			if gcd(m, n) != 1 {
				continue
			}

			// Generate a, b, c
			a := (m * m) - (n * n)
			b := 2 * m * n
			c := (m * m) + (n * n)

			// See if a scalar multiple exists such that k(a+b+c) = num
			if num % (a + b + c) == 0 {
				k := num / (a + b + c)
				a *= k
				b *= k
				c *= k

				return a * b * c
			}
		}
	}

	return -1
}

func gcd(a,b int) int {
	var gcdVal int
	for i := 1; i <= a && i <= b ; i++ {
		if(a % i == 0 && b % i == 0) {
			gcdVal = i
		} 
	}
	return gcdVal
}  

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(approach1(num))
}