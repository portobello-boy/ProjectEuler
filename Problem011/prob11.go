package main

import (
	"os"
	"fmt"
	"strconv"
	"strings"
	"math"
	"io/ioutil"
)

func buildGrid(text string) [][]int {
	rows := strings.Split(text, "\n")
	grid := make([][]int, len(rows))

	for rowInd, row := range rows {
		splitRow := strings.Split(row, " ")
		grid[rowInd] = make([]int, len(splitRow))

		for colInd, entry := range splitRow {
			num, _ := strconv.Atoi(entry)
			grid[rowInd][colInd] = num
		}
	}

	return grid
}

func buildTranspose(grid [][]int) [][]int {
	trans := make([][]int, len(grid))
	for i := 0; i < len(trans); i++ {
		trans[i] = make([]int, len(grid[0]))
	}

	for i, row := range grid {
		for j, _ := range row {
			trans[j][i] = grid[i][j]
		}
	}

	return trans
}

func windowProd(window []int) int {
	prod := 1
	for _, num := range window {
		prod *= num
	}
	return prod
}

func largestHorizontalProd(grid [][]int, size int) int {
	largest := 0

	left := 0
	right := size

	if right >= len(grid[0]) {
		fmt.Println("Window size too big for given grid")
		return -1
	}
	
	for _, row := range grid {
		for right <= len(grid[0]) {
			if windowProd(row[left:right]) > largest {
				largest = windowProd(row[left:right])
			}

			left++
			right++
		}
	}

	return largest
}

func largestDiagProd(grid [][]int, size int) int {
	largest := 0

	if size >= len(grid[0]) {
		fmt.Println("Window size too big for given grid")
		return -1
	}
	
	for i := 0; i <= len(grid)-size; i++ {
		for j := 0; j <= len(grid[0])-size; j++ {
			window := make([]int, size)
			for k := 0; k < size; k++ {
				window[k] = grid[i+k][j+k]
			}

			if windowProd(window) > largest {
				largest = windowProd(window)
			}
		}
	}

	for i := 0; i <= len(grid)-size; i++ {
		for j := size-1; j < len(grid[0]); j++ {
			window := make([]int, size)
			for k := 0; k < size; k++ {
				window[k] = grid[i+k][j-k]
			}

			if windowProd(window) > largest {
				largest = windowProd(window)
			}
		}
	}

	return largest
}

func findLargestProdOfLen(grid [][]int, len int) int {
	trans := buildTranspose(grid)
	largestHori := largestHorizontalProd(grid, len)
	largestVert := largestHorizontalProd(trans, len)
	largestDiag := largestDiagProd(grid, len)

	return int(math.Max(math.Max(float64(largestHori), float64(largestVert)), float64(largestDiag)))
}

func main() {
	file := os.Args[1]

	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	grid := buildGrid(string(content))

	fmt.Println(findLargestProdOfLen(grid, 4))
}