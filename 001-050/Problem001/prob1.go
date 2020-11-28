package main

import (
	"os"
	"fmt"
	"strconv"
)

// This is a naive but acceptable approach - iterate through all numbers in range [1, 1000),
// check if they are divisible by 3 or 5, and add them to the total sum if they are.
// Complexity: O(n) for range [1, n)
func approach1(num int) {
	sum := 0

	for i := 1; i < num; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			sum += i
		}
	}

	fmt.Println(sum)
}

// The second approach leverages some number theory and combinatorics. The sum of all multiples
// of 3 would be 3 + 6 + 9 + 12 + ... + 999 = 3 * (1 + 2 + 3 + ... + 333). The second summation
// here (1 + 2 + ... + 333) is a triangular sum which Gauss supposedly found to be n*(n+1)/2. 
// Using this, simple arithmetic can be used to find the values that we so desire. Therefore
// this solution can be calculated in constant time, without the need for any loops.
// Complexity: O(1) for any n
func approach2(num int) {
	fmt.Println(sumOfMults(3, num) + sumOfMults(5, num) - sumOfMults(3 * 5, num))
}

func triangle(val int) int {
	return val * (val+1)/2
}

func sumOfMults(mult int, bound int) int {
	return mult * triangle(bound/mult)
}

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	// approach1(num)
	approach2(num)
}