package factorial

func Factorial(n int64) int64 {
	prod := int64(1)

	for n > 1 {
		prod *= int64(n)
		n -= 1
	}

	return prod
}

func DoubleFactorial(n int64) int64 {
	prod := int64(1)

	for n > 1 {
		prod *= int64(n)
		n -= 2
	}

	return prod
}
