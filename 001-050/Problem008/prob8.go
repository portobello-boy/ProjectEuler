package main

import (
	"os"
	"fmt"
	"strconv"
	"strings"
)

func windowProd(window []string) int {
	prod := 1
	for i := range window {
		num, _ := strconv.Atoi(window[i])
		prod *= num
	}
	return prod
}

func largestProd(num string, size int) int {
	largest := 0

	left := 0
	right := size

	digits := strings.Split(num, "")
	if right >= len(digits) {
		fmt.Println("Window size too big for given number")
		return -1
	}
	
	for right <= len(digits) {
		if windowProd(digits[left:right]) > largest {
			largest = windowProd(digits[left:right])
		}

		left++
		right++
	}

	return largest
}

func main() {
	num := os.Args[1]

	size, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(largestProd(num, size))
}