package modulo

import (
	"errors"
)

func ModularAdder(N int) (func(int) int, error) {
	if N < 1 {
		return nil, errors.New("modulo N must be a Natural number")
	}
	modulo := N
	sum := 0
	return func(n int) int {
		sum += n
		sum %= modulo
		return sum
	}, nil
}