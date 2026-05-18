package main

import (
	"fmt"
	"math"
)

var target int = 2000000

func g(m, n, r, s int) int {
	return (m - r + 1) * (n - s + 1)
}

func f(m, n int) int {
	number_of_rectangles := 0
	for r := range m {
		for s := range n {
			number_of_rectangles += g(m, n, r+1, s+1)
		}
	}
	return number_of_rectangles
}

func main() {
	fmt.Println(f(3, 3))
	minimum_area := 0
	difference_from_target := target
	for m := range 100 {
		for n := range m {
			number_of_rectangles := f(m, n+1)
			if math.Abs(float64(number_of_rectangles-target)) < float64(difference_from_target) {
				minimum_area = m * (n + 1)
				difference_from_target = int(math.Abs(float64(number_of_rectangles - target)))
			}
			if m%10 == 0 && (n+1)%10 == 0 {
				fmt.Println(m, n+1, number_of_rectangles, difference_from_target)
			}
		}
	}
	fmt.Println(minimum_area)

}
