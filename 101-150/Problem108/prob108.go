package main

import (
	"MathLibGo/numberTheory/divisors"
	"fmt"
)

func getCountOfDivisorsOfSquare(divisors []int) int {
	numSet := make(map[int]bool, len(divisors)*len(divisors))
	for _, d1 := range divisors {
		for _, d2 := range divisors {
			numSet[d1*d2] = true
		}
	}

	// fmt.Println(numSet)

	count := 0
	for _, v := range numSet {
		if v {
			count += 1
		}
	}

	return count
}

func main() {
	// divisors, _ := divisors.DivisorList(16)
	// fmt.Println(len(divisors) / 2)

	fmt.Println(getCountOfDivisorsOfSquare([]int{1, 2, 4, 8}))

	for i := 4; ; i++ {
		d, _ := divisors.DivisorList(i)
		numDivisorsOfSquare := getCountOfDivisorsOfSquare(d)

		if numDivisorsOfSquare > 2000 {
			fmt.Printf("%d - %d, %d, %v\n", i, numDivisorsOfSquare, len(d), d)
			break
		}

		if i%100 == 0 {
			fmt.Printf("Checking %v\n", i)
			// fmt.Printf("%d - %d, %d, %v\n", i, numDivisorsOfSquare, len(d), d)

		}
	}
}
