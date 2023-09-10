package primes

import (
	"MathLibGo/constants"
	"container/list"
	"errors"
)

func NextPrime(p int) (int, error) {
	if p < 0 {
		return -1, errors.New("p must be a non-negative integer")
	}
	for {
		if p%2 == 0 {
			p += 1
		} else {
			p += 2
		}
		if IsPrime(p) {
			return p, nil
		}
	}
}

func IsPrime(num int) bool {
	if num%2 == 0 || num%3 == 0 {
		return false
	}

	for k := 1; (6*k-1)*(6*k-1) <= num; k++ {
		if num%(6*k-1) == 0 || num%(6*k+1) == 0 {
			return false
		}
	}

	return true
}

func boundedPrimeSequenceProducer(bound int) <-chan int {
	// Initialize channel for primes
	primes := make(chan int, bound)

	// Create goroutine to produce to channel
	go func() {
		defer close(primes)
		boundReached := false

		// Iterate over first primes stored in constants, produce them to channel
		for _, p := range constants.FirstPrimes {
			if p > bound {
				boundReached = true
				break
			}
			primes <- p
		}

		if boundReached {
			return
		}

		// Start checking candidate numbers for primality
		k := int((constants.FirstPrimes[len(constants.FirstPrimes)-1] + 5) / 6)
		current := 0

		for current < bound {
			current = 6*k - 1

			if current >= bound {
				break
			}

			if IsPrime(current) {
				primes <- current
			}

			current += 2

			if current >= bound {
				break
			}

			if IsPrime(current) {
				primes <- current
			}

			k += 1
		}
	}()

	// Return channel for processing
	return primes
}

func BoundedPrimeSequence(bound int) []int {
	// Allocate initial memory, number of primes is always less than the bound
	primeContainer := make([]int, bound)
	i := 0

	for p := range boundedPrimeSequenceProducer(bound) {
		primeContainer[i] = p
		i++
	}

	return primeContainer[:i]
}

func primeSequenceProducer(n int) <-chan int {
	// Initialize channel for primes
	primes := make(chan int)

	// Create goroutine to produce to channel
	go func() {
		defer close(primes)

		// Initialize count
		count := 0

		// Iterate over first primes stored in constants, produce them to channel
		for _, p := range constants.FirstPrimes {
			if count++; count > n {
				return
			}

			primes <- p
		}

		if count++; count > n {
			return
		}

		// Start checking candidate numbers for primality
		k := int((constants.FirstPrimes[len(constants.FirstPrimes)-1] + 5) / 6)
		current := 0

		for count < n {
			current = 6*k - 1

			if IsPrime(current) {
				primes <- current

				if count++; count > n {
					return
				}
			}

			current += 2

			if IsPrime(current) {
				primes <- current

				if count++; count > n {
					return
				}
			}

			k += 1
		}
	}()

	// Return channel for processing
	return primes
}

func PrimeSequence(n int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("n must be a Natural number")
	}
	// Initialize linked list for primes to avoid repeated copying of array
	primeContainer := list.New()
	count := 0

	// Get primes and add to list, track count
	for p := range primeSequenceProducer(n) {
		primeContainer.PushBack(p)
		count += 1
	}

	// Initialize slize of correct size, add primes to slice
	primes := make([]int, count)
	i := 0
	for p := primeContainer.Front(); p != nil; p = p.Next() {
		primes[i] = p.Value.(int)
		i++
	}

	return primes, nil
}
