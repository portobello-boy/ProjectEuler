package main

import (
	"os"
	"fmt"
	"strconv"
	"math"
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

func sumOverArray(nums []int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}

func generatePrimesInBound(bound int) []int {
	// So we don't allocate needless amounts of memory, we make an array with a known upper
	// bound on pi(x) for x >= 17
	primes := make([]int, int(math.Round(1.25506 * (float64(bound) / math.Log(float64(bound))))))
	primes[0] = 2
	primes[1] = 3

	index := 2			// How many primes have we encountered
	k := 1				// Multiple of 6
	possiblePrime := 3	// What number are we looking at

	for possiblePrime < bound {
		possiblePrime = 6*k - 1
		if possiblePrime >= bound { // Make sure we aren't past the bound
			break
		}

		if isPrime(possiblePrime) {
			primes[index] = possiblePrime
			index++
		}

		possiblePrime += 2
		if possiblePrime >= bound { // This statement is needed in the event that the desired prime is a twin prime
			break
		}

		if isPrime(possiblePrime) {
			primes[index] = possiblePrime
			index++
		}
		k++
	}

	return primes
}

func approach1(bound int) int {
	return sumOverArray(generatePrimesInBound(bound))
}

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(approach1(num))
}