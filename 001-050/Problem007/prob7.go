package main

import (
	"os"
	"fmt"
	"strconv"
)

func isPrime(num int) bool {
	if num % 2 == 0 || num % 3 == 0 {
		return false
	}

	for k := 1; (6*k-1) * (6*k-1) <= num; k++ {
		if num % (6*k - 1) == 0 || num % (6*k + 1) == 0 {
			return false
		}
	}

	return true
}

// All integers can be written as 6*k + [6], where [6] is the modulo set Z/6. Anything
// of the form 6k + 0, 6k + 2, or 6k + 4 is divisible by 2, so therefore composite.
// Anything of the form 6k + 3 is divisible by 3, and thus composite as well. That
// leaves us to check only numbers of the form 6k + 1 and 6k + 5 = 6k - 1, greatly
// improving speed of searching.
func approach1(num int) int {
	if num <= 0 {
		return -1
	} else if num == 1 {
		return 2
	} else if num == 2 {
		return 3
	} else {
		index := 2			// How many primes have we encountered
		k := 1				// Multiple of 6
		possiblePrime := 3	// What number are we looking at

		for index < num {
			possiblePrime = 6*k - 1
			if isPrime(possiblePrime) {
				index++
			}

			if index == num { // This statement is needed in the event that the desired prime is a twin prime
				break
			}

			possiblePrime += 2
			if isPrime(possiblePrime) {
				index++
			}
			k++
		}

		return possiblePrime
	}
}

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(approach1(num))
}