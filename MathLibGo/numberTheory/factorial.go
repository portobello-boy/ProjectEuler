package numberTheory

func DoubleFactorial(n int) int64 {
	prod := int64(1)

	for n > 1 {
		prod *= int64(n)
		n -= 2
	}

	return prod
}
