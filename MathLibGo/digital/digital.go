package digital

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func GetDigitSlice(n int64) ([]int64, error) {
	digitStrs := strings.Split(strconv.Itoa(int(n)), "")
	digits := make([]int64, len(digitStrs))

	for i, d := range digitStrs {
		nd, _ := strconv.Atoi(d)
		digits[i] = int64(nd)
	}
	return digits, nil
}

func SliceToNumber(s []int64) (int64, error) {
	n := int64(0)
	slices.Reverse(s)
	for i, d := range s {
		n += d * int64(math.Pow(10, float64(i)))
	}
	return n, nil
}
