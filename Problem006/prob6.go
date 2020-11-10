package main

import (
	"os"
	"fmt"
	"strconv"
)

func sumOfSqr(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i*i
	}
	return sum
}

func sqrOfSum(num int) int {
	sum := (num * (num+1))/2
	return sum * sum
}

func sumSqrDiff(num int) int {
	return sqrOfSum(num) - sumOfSqr(num)
}

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(sumSqrDiff(num))
}