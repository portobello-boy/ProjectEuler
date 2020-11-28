package main

import (
	"os"
	"fmt"
	"strconv"
	"math"
)

// Given a number of digits, find the largest palindrome by multiplying all numbers
// in descending order
func approach1(num float64) {
	var max float64 = 0
	for i := math.Pow(10, num)-1; i > 99; i-- {
		for j := math.Pow(10, num)-1; j >= i; j-- {
			if isPalindrome(int(i*j)) && i*j > max {
				max = i*j
			}
		}
	}
	fmt.Println(max)
}

func reverse(str string) (result string) {
	for _, c := range str {
		result = string(c) + result
	}
	return
}

func isPalindrome(num int) bool {
	return strconv.Itoa(num) == reverse(strconv.Itoa(num))
}

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	approach1(float64(num))
}