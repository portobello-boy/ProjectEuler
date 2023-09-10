package main

import (
	"fmt"
	"math"
	"numberTheory"
	"time"
)

var exploredTotients map[int]int = make(map[int]int)
var primes []int = make([]int, 0)
var bound int = 10000000

func getCharCountMap(str string) map[rune]int {
	charCount := make(map[rune]int)
	for _, char := range str {
		charCount[char]++
	}
	return charCount
}

func checkPermutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1CharCountMap := getCharCountMap(s1)
	s2CharCountMap := getCharCountMap(s2)
	if len(s1CharCountMap) != len(s2CharCountMap) {
		return false
	}
	for s1Char, s1Count := range s1CharCountMap {
		if s1Count != s2CharCountMap[s1Char] {
			return false
		}
	}
	return true
}

func getPrimeFactors(num int) map[int]int {
	m := make(map[int]int)

	for _, p := range primes {
		for num%p == 0 {
			num = num / p
			m[p]++
		}
		if num == 1 {
			return m
		}
	}
	return nil
}

func totient(a int) int {
	totient := 0
	if a > 4 && a%4 == 0 {
		totient = 2 * exploredTotients[a/2]
	} else if a > 4 && a%2 == 0 {
		totient = exploredTotients[a/2]
	} else {
		prod := float64(a)
		for p := range getPrimeFactors(a) {
			prod *= (1.0 - 1.0/float64(p))
		}
		totient = int(math.Round(prod))
	}
	exploredTotients[a] = totient
	return totient
}

func main() {
	t := time.Now()
	out := make(chan int)
	go numberTheory.SieveOfEratosthenes(bound, out)
	for p := range out {
		primes = append(primes, p)
	}

	minN := 2
	minFunc := float64(minN) / float64(totient(minN))

	for i := 2; i <= bound; i++ {
		tot := totient(i)
		if !checkPermutation(fmt.Sprint(i), fmt.Sprint(tot)) {
			continue
		}
		fmt.Println(i, tot, "time so far:", time.Since(t))
		totFunc := float64(i) / float64(tot)
		if totFunc < minFunc {
			minN = i
			minFunc = totFunc
		}
	}

	fmt.Println(minN, minFunc)

	fmt.Println("Took", time.Since(t))
}
