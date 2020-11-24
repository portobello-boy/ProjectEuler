package main

import (
	"fmt"
	"time"
)

func main() {
	cur := time.Date(1901, 1,  1,  0, 0, 0, 0, time.UTC)
	end := time.Date(2000, 12, 31, 0, 0, 0, 0, time.UTC)
	sundayCount := 0

	for cur.Before(end) {
		cur = cur.AddDate(0, 1, 0)
		if cur.Day() == 1 && cur.Weekday().String() == "Sunday" {
			sundayCount++
		}
	}

	fmt.Println(sundayCount)
}