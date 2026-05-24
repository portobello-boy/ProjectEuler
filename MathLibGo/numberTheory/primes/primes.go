package primes

import (
	"MathLibGo/constants"
	"container/list"
	"errors"
	"math"
	"math/big"
	"sync"
)

var MillerRabinPrimarityCache = map[uint64]bool{}
var primeLock = sync.Mutex{}

func Lcm(a, b int) int {
	return int(math.Abs(float64(a)) / float64(Gcd(a, b)) * math.Abs(float64(b)))
}

func Gcd(a, b int) int {

	// Set a and b
	a = int(math.Abs(float64(a)))
	b = int(math.Abs(float64(b)))

	if a == b {
		return a
	} else if a == 0 {
		return b
	} else if b == 0 {
		return a
	} else if a&1 == 0 { // a is even
		if b&1 == 0 { // b is even
			return 2 * Gcd(a>>1, b>>1)
		}
		return Gcd(a>>1, b) // b is odd
	}
	if b&1 == 0 {
		return Gcd(a, b>>1)
	} else if a > b && b&1 != 0 {
		return Gcd((a-b)>>1, b)
	}
	return Gcd((b-a)>>1, a)
}

func NextPrime(p uint) (uint, error) {
	if p < 0 {
		return 0, errors.New("p must be a non-negative integer")
	}
	if p == 1 {
		return 2, nil
	}
	if p == 2 {
		return 3, nil
	}
	for {
		if p%2 == 0 {
			p += 1
		} else {
			p += 2
		}
		if IsProbablePrime(uint64(p)) {
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

func IsProbablePrime(num uint64) bool {
	primeLock.Lock()
	defer primeLock.Unlock()

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
	bases := []uint64{2, 325, 9375, 28178, 450775, 9780504, 1795265022}

	check := func(base uint64) bool {
		// Need big int to help calculate
		bigM := new(big.Int).Exp(big.NewInt(int64(base)), big.NewInt(int64(d)), big.NewInt(int64(num)))

		m := uint64(bigM.Uint64())

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

func PollardRhoBrentCycleFactor(num uint64) uint64 {
	if num%2 == 0 {
		return 2
	}
	if num%3 == 0 {
		return 3
	}

	for _, c := range []uint64{1, 3, 5, 7, 11, 13, 17, 19, 23} {
		f := func(x uint64) uint64 {
			return ((x * x) + c) % num
		}

		x := uint64(2)
		y := x
		d := uint64(1)

		for d == 1 {
			x = f(x)
			y = f(f(y))
			d = uint64(Gcd(int(math.Abs(float64(x-y))), int(num)))
		}

		if d == num {
			return num
		} else {
			return d
		}

	}
	return num

	/*
			    # Deterministic parameter choices
		    # Brent cycle detection with polynomial f(x) = x^2 + c mod n
		    for c in (1, 3, 5, 7, 11, 13, 17, 19, 23):
		        y = 2
		        m = 128
		        g = 1
		        r = 1
		        q = 1

		        def f(x: int) -> int:
		            return (x * x + c) % n

		        while g == 1:
		            x = y
		            for _ in range(r):
		                y = f(y)
		            k = 0
		            while k < r and g == 1:
		                ys = y
		                for _ in range(min(m, r - k)):
		                    y = f(y)
		                    q = (q * abs(x - y)) % n
		                g = _gcd(q, n)
		                k += m
		            r <<= 1

		        if g == n:
		            # fallback: standard gcd steps
		            g = 1
		            y = ys
		            while g == 1:
		                y = f(y)
		                g = _gcd(abs(x - y), n)

		        if 1 < g < n:
		            return g
			return nun
	*/
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
