package main

import (
	"os"
	"fmt"
	"strconv"
)

// Approach 1 uses a traditional brute force method - while the current triangular
// number has less than the desired number of divisors (quickly determined using up
// to the sqrt(n)), then generate more.
func approach1(num int) int {
	ind := 1
	for getDivisorCountBrute(getTriangleNum(ind)) < num {
		ind++
	}
	return getTriangleNum(ind)
}

// Approach 2 uses a more sophisticated approach with the Tau function. Given a prime
// factorization for n = (p^a)(q^b)(r^c)... then the number of divisors d(n)=(a+1)(b+1)(c+1)...
func approach2(num int) int {
	ind := 1
	for getDivisorCountTau(getTriangleNum(ind)) < num {
		ind++
	}
	return getTriangleNum(ind)
}

func getTriangleNum(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	return sum
}

func getDivisorCountBrute(num int) int {
	count := 0
	for i := 1; i*i <= num; i++ {
		if num % i == 0 {
			count += 2
		}
	}
	return count
}

func getDivisorCountTau(num int) int {
	factors := getPrimeFactors(num)
	prod := 1
	for _, val := range factors {
		prod *= (val + 1)
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

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(approach2(num))
}