package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func printGrid(grid [][]uint8) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func concatDigits(digits ...uint8) int {
	totalStr := ""
	for _, digit := range digits {
		totalStr += strconv.FormatUint(uint64(digit), 10)
	}
	total, err := strconv.Atoi(totalStr)
	check(err)
	return total
}

func isValidEntry(grid [][]uint8, num uint8, rowIndex, colIndex int) bool {
	// Check row
	if slices.Contains[[]uint8, uint8](grid[rowIndex], num) {
		return false
	}

	// Check column
	for _, row := range grid {
		if num == row[colIndex] {
			return false
		}
	}

	// Check box
	for _, row := range grid[3*(rowIndex/3) : 3*(rowIndex/3)+3] {
		if slices.Contains[[]uint8, uint8](row[3*(colIndex/3):3*(colIndex/3)+3], num) {
			return false
		}
	}

	return true
}

func getValidEntries(grid [][]uint8, rowIndex, colIndex int) []uint8 {
	validNumbers := []uint8{}

	for num := uint8(1); num < 10; num++ {
		if isValidEntry(grid, num, rowIndex, colIndex) {
			validNumbers = append(validNumbers, num)
		}
	}

	return validNumbers
}

func findEmptySpot(grid [][]uint8) (int, int) {
	for rowIndex, row := range grid {
		for colIndex, entry := range row {
			if entry == 0 {
				return rowIndex, colIndex
			}
		}
	}

	return -1, -1
}

func solveSinglePuzzle(grid [][]uint8) ([][]uint8, error) {
	// Find first empty cell
	rowIndex, colIndex := findEmptySpot(grid)

	if colIndex == rowIndex && rowIndex == -1 {
		return grid, nil
	}

	// Get valid entries
	validNumbers := getValidEntries(grid, rowIndex, colIndex)

	// If there are no valid entries, break from this tree
	if len(validNumbers) == 0 {
		return grid, errors.New("NoValidSolution")
	}

	// For each valid number, asign it to the grid and recursively solve
	for _, num := range validNumbers {
		grid[rowIndex][colIndex] = num

		grid, err := solveSinglePuzzle(grid)

		if err == nil {
			return grid, nil
		}
		grid[rowIndex][colIndex] = 0
	}

	return grid, errors.New("NoValidSolution")
}

func loadSinglePuzzle(scanner *bufio.Scanner) [][]uint8 {

	grid := make([][]uint8, 9)
	for i := range grid {
		grid[i] = make([]uint8, 9)
	}

	for i := 0; i < 9; i++ {
		str := scanner.Text()

		str = strings.TrimSuffix(str, "\n")

		rowStr := strings.Split(str, "")
		for j, digitStr := range rowStr {
			digit, err := strconv.Atoi(digitStr)
			check(err)

			grid[i][j] = uint8(digit)
		}

		scanner.Scan()
	}

	return grid
}

func solveProblem(filename string) int {

	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()

	sum := 0

	for scanner.Scan() {
		grid := loadSinglePuzzle(scanner)
		grid, err := solveSinglePuzzle(grid)
		check(err)
		sum += concatDigits(grid[0][0:3]...)
	}

	return sum
}

func main() {

	fmt.Println(solveProblem("./puzzles.txt"))
}
