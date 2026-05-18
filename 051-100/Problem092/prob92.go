package main

import (
	helper "MathLibGo"
	"MathLibGo/digital"
	"fmt"
)

const max = 10000000

func main() {
	cached_squareDigits := helper.CacheWrapper_SingleToSingle(func(n int) int {
		digits, _ := digital.GetDigitSlice(int64(n))
		sum := 0
		for _, d := range digits {
			sum += int(d * d)
		}
		return sum
	})

	count := 0

	for i := 1; i < max; i++ {
		n := i
		for n != 89 && n != 1 {
			n = cached_squareDigits(n)
		}
		if n == 89 {
			count++
		}
	}

	fmt.Println(count)
}
