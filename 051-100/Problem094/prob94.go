package main

import (
	"fmt"
	"math"
)

const MAX_PERIMETER = 1000000000
const BOUND = MAX_PERIMETER / 3

func isIntegerArea(height float64, base int) bool {
	area := 0.5 * float64(base) * height
	return area == math.Trunc(area)
}

func getIsoscelesHeight(a, b int) float64 {
	return math.Sqrt(math.Pow(float64(b), 2) - 0.25*math.Pow(float64(a), 2))
}

func prob94(bound int) int {
	perimeterSum := 0 // n=1 case

	for n := 1; n <= bound; n += 2 {
		if (n+1)%2 == 0 {
			// fmt.Println("n, n+1", n, n+1)
			val1 := math.Sqrt(float64(((3 * n) + 1) * (n - 1)))
			// fmt.Println(val1)
			// fmt.Println(val1 * (float64(n+1) / 4.0))
			if val1 == math.Trunc(val1) {
				perimeterSum += 3*n + 1
			}
		}
		if (n-1)%2 == 0 {
			// fmt.Println("n, n-1", n, n-1)
			val2 := math.Sqrt(float64(((3 * n) - 1) * (n + 1)))
			// fmt.Println(val2)
			if val2 == math.Trunc(val2) {
				perimeterSum += 3*n - 1
			}
		}

		// fmt.Println(n, n-1, getIsoscelesHeight(n-1, n))
		// if isIntegerArea(getIsoscelesHeight(n-1, n), n-1) {
		// 	perimeterSum += 3*n - 1
		// }
		// fmt.Println(n, n+1, getIsoscelesHeight(n+1, n))
		// if isIntegerArea(getIsoscelesHeight(n+1, n), n+1) {
		// 	perimeterSum += 3*n + 1
		// }
	}
	return perimeterSum
}

func main() {
	// fmt.Println("Hello World")
	// fmt.Println(MAX_PERIMETER, BOUND)
	// fmt.Println(isIntegerArea(getIsoscelesHeight(6, 5), 6))
	fmt.Println(prob94(BOUND))
}
