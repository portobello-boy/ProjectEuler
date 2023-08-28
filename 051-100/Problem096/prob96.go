package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isValidPuzzle(grid [][]uint8) bool {
	return true
}

func solveSinglePuzzle(grid [][]uint8) ([][]uint8, error) {
	return nil, nil
}

func loadSinglePuzzle(filename string) [][]uint8 {

	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	reader := bufio.NewReader(f)

	_, err = reader.ReadString('\n')
	check(err)

	grid := make([][]uint8, 9)
	for i := range grid {
		grid[i] = make([]uint8, 9)
	}

	for i := 0; i < 9; i++ {
		str, err := reader.ReadString('\n')
		check(err)

		str = strings.TrimSuffix(str, "\n")

		rowStr := strings.Split(str, "")
		for j, digitStr := range rowStr {
			digit, err := strconv.Atoi(digitStr)
			check(err)

			grid[i][j] = uint8(digit)
		}

		// fmt.Println(str)
	}

	return grid
}

func main() {
	fmt.Println(loadSinglePuzzle("./single.txt"))
}
