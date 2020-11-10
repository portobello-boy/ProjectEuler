package main

import (
	"os"
	"fmt"
	"strconv"
	"math"
)

func mapUnion(insert map[int]int, other map[int]int) {
	for key, val := range other {
		if v, ok := insert[key]; ok == false {
			insert[key] = val
		} else if val > v {
			insert[key] = val
		}
	}
}

func multOverMap(m map[int]int) int {
	prod := 1
	for key, val := range m {
		prod *= int(math.Pow(float64(key), float64(val)))
	}
	return prod
}

func getPrimeFactors(num int) map[int]int {
	m := make(map[int]int)
	for num % 2 == 0 {
		m[2]++
		num /= 2
	}

	for i := 3; i*i <= num; i += 2 {
		for num % i == 0 {
			m[i]++
			num /= i
		}
	}

	if num > 2 {
		m[num]++
	}

	return m
}

func lowestCommonMultiple(num int) int {
	m := make(map[int]int)
	for i := 2; i <= num; i++ {
		mapUnion(m, getPrimeFactors(i))
	}

	return multOverMap(m)
}

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(lowestCommonMultiple(num))
}