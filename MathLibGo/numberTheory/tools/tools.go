package tools

import "math"

func SliceSumInt(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func HypotenuseInt(a, b int) int {
	return int(math.Sqrt(math.Pow(float64(a), 2) + math.Pow(float64(b), 2)))
}
