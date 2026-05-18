package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Priority Queue implementation from: https://pkg.go.dev/container/heap#example-package-PriorityQueue
// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	row      int
	col      int
	path     []int
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest, not hight, priority so we use less than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadRow(scanner *bufio.Scanner) []int {
	str := scanner.Text()
	str = strings.TrimSuffix(str, "\n")
	rowStr := strings.Split(str, ",")

	row := make([]int, len(rowStr))

	for i, val := range rowStr {
		digit, _ := strconv.Atoi(val)
		row[i] = digit
	}

	return row
}

func readFile(filename string) [][]int {

	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// scanner.Scan()

	grid := make([][]int, 0)

	for scanner.Scan() {
		grid = append(grid, loadRow(scanner))
	}

	return grid
}

func findMin(v []int) int {
	m := 0
	for i, e := range v {
		if i == 0 || e < m {
			m = e
		}
	}
	return m
}

func main() {
	grid := readFile("./matrix.txt")
	grid_rows := len(grid)
	grid_cols := len(grid[0])
	fmt.Println(grid)
	fmt.Println(grid_rows, grid_cols)

	min_vals := make([]int, 0)
	min_path_sum := make([]int, 0)
	explored := make(map[string]bool)

	pq := make(PriorityQueue, len(grid))

	for i := range grid {
		// fmt.Println(i, start)

		start_value := grid[i][0]
		pq[i] = &Item{
			value:    start_value,
			row:      i,
			col:      0,
			path:     []int{start_value},
			priority: start_value,
			index:    i,
		}

		explored_key := fmt.Sprintf("%d,%d,%d", i, 0, start_value)
		explored[explored_key] = true
	}

	heap.Init(&pq)

	for pq.Len() > 0 {
		// fmt.Println("working...")
		item := heap.Pop(&pq).(*Item)

		current_row := item.row
		current_col := item.col

		explored_key := fmt.Sprintf("%d,%d,%d", current_row, current_col, item.value)
		explored[explored_key] = true
		// fmt.Println(explored_key)

		// fmt.Println(item.value, current_row, current_col)

		if current_col == grid_cols-1 {
			min_path_sum = append(min_path_sum, item.priority)
		}

		if current_row-1 >= 0 {
			new_value := grid[current_row-1][current_col]

			check_key := fmt.Sprintf("%d,%d,%d", current_row-1, current_col, new_value)
			if _, ok := explored[check_key]; !ok {
				heap.Push(&pq, &Item{
					value:    new_value,
					row:      current_row - 1,
					col:      current_col,
					path:     append(item.path, new_value),
					priority: item.priority + new_value,
					index:    0,
				})
			}
		}
		if current_row+1 < grid_rows {
			new_value := grid[current_row+1][current_col]

			check_key := fmt.Sprintf("%d,%d,%d", current_row+1, current_col, new_value)
			if _, ok := explored[check_key]; !ok {
				heap.Push(&pq, &Item{
					value:    new_value,
					row:      current_row + 1,
					col:      current_col,
					path:     append(item.path, new_value),
					priority: item.priority + new_value,
					index:    0,
				})
			}
		}
		if current_col+1 < grid_cols {
			new_value := grid[current_row][current_col+1]

			check_key := fmt.Sprintf("%d,%d,%d", current_row, current_col+1, new_value)
			if _, ok := explored[check_key]; !ok {
				heap.Push(&pq, &Item{
					value:    new_value,
					row:      current_row,
					col:      current_col + 1,
					path:     append(item.path, new_value),
					priority: item.priority + new_value,
					index:    0,
				})
			}
		}
	}

	min_vals = append(min_vals, findMin(min_path_sum))

	fmt.Println(findMin(min_vals))

}
