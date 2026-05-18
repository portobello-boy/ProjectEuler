package main

import (
	"MathLibGo/digital"
	"MathLibGo/operations/factorial"
	"fmt"
)

func GetDigitalFactorialChain(n int64) int {
	digitalFactorialMap := make(map[int64]int64)

	sum := int64(0)
	for _, ok := digitalFactorialMap[n]; !ok; _, ok = digitalFactorialMap[n] {
		ds, _ := digital.GetDigitSlice(n)
		sum = int64(0)
		for _, d := range ds {
			sum += factorial.Factorial(d)
		}

		digitalFactorialMap[n] = sum
		n = sum
	}

	return len(digitalFactorialMap)
}

func main() {
	count := 0
	for n := int64(1); n < 1000000; n++ {
		if GetDigitalFactorialChain(n) == 60 {
			count += 1
		}
	}
	fmt.Println(count)
}
