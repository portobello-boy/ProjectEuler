package numberproperties

import (
	"MathLibGo/numberTheory/divisors"
	"MathLibGo/numberTheory/tools"
	"errors"
)

func Totient(n int) (int, error) {
	if n < 1 {
		return -1, errors.New("n must be a Natural number")
	}
	primeMap, _ := divisors.PrimeDivisorMap(n)

	product := float64(n)
	for k := range primeMap {
		product *= (1.0 - 1.0/float64(k))
	}

	return int(product), nil
}

func IsPerfect(n int) (bool, error) {
	if n < 1 {
		return false, errors.New("n must be a Natural number")
	}

	properDivisors, _ := divisors.ProperDivisorList(n)
	sum := tools.SliceSumInt(properDivisors...)

	return sum == n, nil
}

func IsAmicable(n int) (bool, error) {
	if n < 1 {
		return false, errors.New("n must be a Natural number")
	}

	properDivisorsN, _ := divisors.ProperDivisorList(n)
	m := tools.SliceSumInt(properDivisorsN...)

	properDivisorsM, _ := divisors.ProperDivisorList(m)
	p := tools.SliceSumInt(properDivisorsM...)

	return p == n && m != n, nil
}

func IsAmicablePair(m, n int) (bool, error) {
	if n < 1 || m < 1 {
		return false, errors.New("m and n must be a Natural numbers")
	}

	properDivisorsN, _ := divisors.ProperDivisorList(n)
	p := tools.SliceSumInt(properDivisorsN...)

	return p == m && m != n, nil
}

func IsAbundant(n int) (bool, error) {
	if n < 1 {
		return false, errors.New("n must be a Natural number")
	}

	properDivisors, _ := divisors.ProperDivisorList(n)
	sum := tools.SliceSumInt(properDivisors...)

	return sum > n, nil
}

func IsDeficient(n int) (bool, error) {
	if n < 1 {
		return false, errors.New("n must be a Natural number")
	}

	properDivisors, _ := divisors.ProperDivisorList(n)
	sum := tools.SliceSumInt(properDivisors...)

	return sum < n, nil
}
