package divisors

import (
	"container/list"
	"errors"
	"math"
	"numberTheory/primes"
)

func divisorListProducer(n int) <-chan int {
	divisors := make(chan int, n)
	divisorsList := list.New()

	go func() {
		defer close(divisors)

		for i := 1; i < int(math.Sqrt(float64(n)))+1; i++ {
			if n%i == 0 {
				divisors <- i
				divisorsList.PushBack(n / i)
			}
		}

		for d := divisorsList.Back(); d != nil; d = d.Prev() {
			divisors <- d.Value.(int)
		}
	}()

	return divisors
}

func DivisorList(n int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("n must be a Natural number")
	}

	divisors := make([]int, n)
	index := 0

	for d := range divisorListProducer(n) {
		divisors[index] = d
		index++
	}

	return divisors[:index], nil
}

func ProperDivisorList(n int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("n must be a Natural number")
	}

	divisors, _ := DivisorList(n)

	return divisors[:len(divisors)-1], nil
}

func PrimeDivisorMap(n int) (map[int]int, error) {
	if n < 0 {
		return nil, errors.New("n must be a Natural number")
	}
	primeFactorMap := make(map[int]int)

	primeList := primes.BoundedPrimeSequence(n)
	primeIndex := 0

	// If n is prime, it will be in the prime list at the end
	if primeList[len(primeList)-1] == n {
		return map[int]int{n: 1}, nil
	}

	for n != 1 {
		prime := primeList[primeIndex]
		for n%prime == 0 {
			n /= prime
			// if _, ok := primeFactorMap[prime]; !ok {
			// 	primeFactorMap[prime] = 0
			// }
			primeFactorMap[prime]++
		}
		primeIndex++
	}

	return primeFactorMap, nil
}
