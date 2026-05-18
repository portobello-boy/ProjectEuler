package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("./base_exp.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	greatestValue, greatestRow, row := 0.0, 0, 0

	for scanner.Scan() {
		str := scanner.Text()
		str = strings.TrimSuffix(str, "\n")
		rowStr := strings.Split(str, ",")

		exp := make([]float64, len(rowStr))

		for i, val := range rowStr {
			digit, _ := strconv.Atoi(val)
			exp[i] = float64(digit)
		}

		row += 1
		value := exp[1] * math.Log(exp[0])

		if value > greatestValue {
			greatestValue = value
			greatestRow = row
		}
	}

	fmt.Println(greatestRow, greatestValue)

}
