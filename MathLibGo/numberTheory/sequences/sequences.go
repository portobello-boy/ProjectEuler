package sequences

import (
	"container/list"
	"errors"
	"math"

	"MathLibGo/constants"
)

func Fibonacci(n int) (int, error) {
	if n < 0 {
		return -1, errors.New("n must be a Natural number")
	}
	return int(math.Round(math.Pow(constants.GoldenRatio, float64(n)) / math.Sqrt(5))), nil
}

func FibonacciSequence(length int) ([]int, error) {
	if length <= 0 {
		return nil, errors.New("length of sequence must be positive")
	}
	seq := make([]int, length)
	for i := 0; i < length; i++ {
		iFib, _ := Fibonacci(i)
		seq[i] = iFib
	}
	return seq, nil
}

func BoundedFibonacciSequence(bound int) ([]int, error) {
	if bound < 0 {
		return nil, errors.New("bound must be non-negative Integer")
	}
	fibList := list.New()
	count := 0

	for i, _ := Fibonacci(count); i < bound; i, _ = Fibonacci(count) {
		fibList.PushBack(i)
		count++
	}

	boundedFibList := make([]int, count)
	index := 0

	for f := fibList.Front(); f != nil; f = f.Next() {
		boundedFibList[index] = f.Value.(int)
		index++
	}

	return boundedFibList, nil
}

func FibonacciClosure() func() int {
	i := 0
	j := 1
	return func() int {
		k := i + j
		i, j = j, k
		return k
	}
}
