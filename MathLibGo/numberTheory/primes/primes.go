package primes

import (
	"MathLibGo/constants"
	"container/list"
	"errors"
	"math/big"
)

var MillerRabinPrimarityCache = map[uint]bool{}

func NextPrime(p uint) (uint, error) {
	if p < 0 {
		return 0, errors.New("p must be a non-negative integer")
	}
	if p == 1 {
		return 2, nil
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

func IsPrime(num uint) bool {
	if num == 2 || num == 3 {
		return true
	}
	if num%2 == 0 || num%3 == 0 {
		return false
	}

	for k := 1; uint(6*k-1)*uint(6*k-1) <= num; k++ {
		if num%uint(6*k-1) == 0 || num%uint(6*k+1) == 0 {
			return false
		}
	}

	return true
}

func IsProbablePrime(num uint) bool {
	if res, ok := MillerRabinPrimarityCache[num]; ok {
		return res
	}

	if num < 2 {
		MillerRabinPrimarityCache[num] = false
		return false
	}

	if num == 2 || num == 3 {
		MillerRabinPrimarityCache[num] = true
		return true
	}

	if num%2 == 0 {
		MillerRabinPrimarityCache[num] = false
		return false
	}
	// Write n-1 as 2^r * d
	d := num - 1
	r := 0
	for d%2 == 0 {
		d /= 2
		r++
	}

	// Test against these SPRP bases sets
	// https://miller-rabin.appspot.com/
	bases := []uint{2, 325, 9375, 28178, 450775, 9780504, 1795265022}

	check := func(base uint) bool {
		// Need big int to help calculate
		bigM := new(big.Int).Exp(big.NewInt(int64(base)), big.NewInt(int64(d)), big.NewInt(int64(num)))

		m := uint(bigM.Uint64())

		if m == 1 || m == (num-1) {
			return true
		}

		for range r - 1 {
			m = (m * m) % num
			if m == num-1 {
				return true
			}
		}
		return false
	}

	for _, base := range bases {
		if base%num == 0 {
			continue
		}
		if !check(base) {
			MillerRabinPrimarityCache[num] = false
			return false
		}
	}

	MillerRabinPrimarityCache[num] = true
	return true

}

func boundedPrimeSequenceProducer(bound uint) <-chan uint {
	// Initialize channel for primes
	primes := make(chan uint, bound)

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
		current := uint(0)

		for current < bound {
			current = uint(6*k - 1)

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

func BoundedPrimeSequence(bound uint) []uint {
	// Allocate initial memory, number of primes is always less than the bound
	primeContainer := make([]uint, bound)
	i := 0

	for p := range boundedPrimeSequenceProducer(bound) {
		primeContainer[i] = p
		i++
	}

	return primeContainer[:i]
}

func primeSequenceProducer(n uint) <-chan uint {
	// Initialize channel for primes
	primes := make(chan uint)

	// Create goroutine to produce to channel
	go func() {
		defer close(primes)

		// Initialize count
		count := uint(0)

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
		current := uint(0)

		for count < n {
			current = uint(6*k - 1)

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

func PrimeSequence(n uint) ([]uint, error) {
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
	primes := make([]uint, count)
	i := 0
	for p := primeContainer.Front(); p != nil; p = p.Next() {
		primes[i] = p.Value.(uint)
		i++
	}

	return primes, nil
}
