package numberproperties

import (
	"errors"
	"math"
	"numberTheory/divisors"
	"numberTheory/tools"
)

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

func Lcm(a, b int) int {
	return int(math.Abs(float64(a)) / float64(Gcd(a, b)) * math.Abs(float64(b)))
}

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
