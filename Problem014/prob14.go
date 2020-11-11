package main

import (
	"os"
	"fmt"
	"strconv"
)

var cache = make(map[int]int)

func collatz(num int) func() int {
	val := num
	return func() int {
		if val == 1 {
			return 0
		} else if val % 2 == 0 {
			val /= 2
		} else {
			val = (3*val) + 1
		}
		return val
	}
}

func getSequenceLength(num int) int {
	nextTerm := collatz(num)
	l := 1

	n := nextTerm()
	for n != 0 {
		if cache[n] != 0 {
			l += cache[n]
			break
		}

		l++
		n = nextTerm()
	}

	cache[num] = l
	return l
}

func getLongestSequence(bound int) (int, int) {
	num := 0
	length := 0

	for i := 1; i < bound; i++ {
		l := getSequenceLength(i)
		if l > length {
			num = i
			length = l
		}
	}

	return num, length
}

func main() {
	bound, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	num, length := getLongestSequence(bound)
	fmt.Printf("%d - length: %d\n", num, length)
}