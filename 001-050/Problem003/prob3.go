package main

import (
	"os"
	"fmt"
	"strconv"
)

// For this approach, I wanted to use Golang closures again. For finding the largest divisor, this
// is unnecessary, but I thought this would be a neat way to individually grab *all* divisors, if
// I so choose.
func approach1(num int) {
	largest := -1
	for num % 2 == 0 {
		largest = 2
		num >>= 1 // left shift 2 is equivalent to (and faster than) /= 2
	}

	nextDivisor := getDivisor(num) // Initialize closure function

	largest = nextDivisor() // Grab the next divisor
	for num != 1 {
		num /= largest // Divide num by the largest divisor, and grab the next one
		largest = nextDivisor()
	}
	fmt.Println(largest)
}

func getDivisor(num int) func() int {
	val := num
	return func() int {
		for i := 3; i*i <= val; i += 2 {
			if val % i == 0 {
				val /= i
				return i
			}
		}
		return val // We are done
	}
}

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	approach1(num)
}