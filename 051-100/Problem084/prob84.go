package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

var dice_size int = 4
var roll_odds float64 = 1.0 / (float64(dice_size * dice_size))
var card_odds float64 = 1.0 / 16.0
var jail int = 10
var start int = 0

var special_squares = map[int]string{
	2: "CC1",
	// 4:  "T1",
	// 5:  "R1",
	7: "CH1",
	// 12: "U1",
	// 15: "R2",
	17: "CC2",
	22: "CH2",
	// 25: "R3",
	// 28: "U2",
	30: "G2J",
	33: "CC3",
	36: "CH3",
	// 38: "T2",
}

func multiplyMatrices(A, B [][]float64) [][]float64 {
	rowsA, colsA := len(A), len(A[0])
	rowsB, colsB := len(B), len(B[0])

	// Check if matrices can be multiplied
	if colsA != rowsB {
		fmt.Println("Matrix multiplication is not valid.")
		return nil
	}

	// Initialize result matrix C
	C := make([][]float64, rowsA)
	for i := range C {
		C[i] = make([]float64, colsB)
	}

	// Perform multiplication
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}

	return C
}

func main() {
	fmt.Println(roll_odds)
	fmt.Println(card_odds * roll_odds)

	probability_map := make([][]float64, 40)
	// Iterate from 0 to 39
	for i := range 40 {
		probability_map[i] = make([]float64, 40)
		if i == 30 {
			continue
		}
		// Iterate over 2 dice rolls
		for j := range dice_size {
			for k := range dice_size {
				// fmt.Println(i, j+1, k+1)
				target_space := (i + j + k + 2) % 40

				// if target_space is a chance or community chest square, add the odds of landing on that square to the
				// target space's odds.
				if id, ok := special_squares[target_space]; !ok {
					if j != k {
						// Add the roll odds to the result of the dice sum
						probability_map[i][target_space] += roll_odds
					} else {
						// Handle double roll
						probability_map[i][target_space] += roll_odds - math.Pow(roll_odds, 3)
						probability_map[i][jail] += math.Pow(roll_odds, 3)
					}
				} else {
					if strings.HasPrefix(id, "CC") {
						probability_map[i][target_space] += (roll_odds * 14.0 * card_odds) // Stay on chance
						probability_map[i][jail] += (roll_odds * card_odds)                // Go to jail
						probability_map[i][start] += (roll_odds * card_odds)               // Advance to Go
					} else if strings.HasPrefix(id, "CH") {
						probability_map[i][target_space] += (roll_odds * 6.0 * card_odds) // Stay on chest
						probability_map[i][jail] += (roll_odds * card_odds)               // Go to jail
						probability_map[i][start] += (roll_odds * card_odds)              // Advance to Go
						probability_map[i][11] += (roll_odds * card_odds)                 // C1
						probability_map[i][24] += (roll_odds * card_odds)                 // E3
						probability_map[i][39] += (roll_odds * card_odds)                 // H2
						probability_map[i][5] += (roll_odds * card_odds)                  // R1

						// Handle back 3 spaces
						back_three := (target_space - 3) % 40
						if back_three == 30 {
							back_three = jail
						}
						probability_map[i][back_three] += (roll_odds * card_odds) // Back 3 spaces

						// Get next railroad from target_space
						next_railroad := int(math.Round(10*math.Ceil(float64(i-5)/10)+5)) % 40
						probability_map[i][next_railroad] += (2.0 * roll_odds * card_odds)

						// Get next utility from target_space
						next_utility := 0
						if target_space < 13 || target_space > 28 {
							next_utility = 12
						} else {
							next_utility = 28
						}
						probability_map[i][next_utility] += (roll_odds * card_odds)
					} else {
						// Landing on G2J
						probability_map[i][jail] += roll_odds
					}
				}

				// fmt.Println((i+j+k+2)%40, probability_map[i][(i+j+k+2)%40])
			}
		}

	}
	for i, k := range probability_map {
		fmt.Println(i, k)
	}

	iterated_matrix := probability_map
	for range 150 {
		iterated_matrix = multiplyMatrices(iterated_matrix, probability_map)
	}
	for i, k := range iterated_matrix {
		fmt.Println(i, k)
	}

	// Iterate over the map and print out the probabilities
	for i, v := range iterated_matrix {
		total := float64(0)
		// fmt.Println(i, iterated_matrix[i])
		for _, p := range v {
			// fmt.Printf("i: %d, ind: %d, p: %.2f\n", i, ind, p)
			total += p

		}
		fmt.Printf("Total for i: %d is %.2f\n", i, total)
	}

	// Get the sum of each possible board space, sorted by highest probability descending
	// This will give us the most likely board space to land on
	sums := make([]float64, 40)
	for _, v := range iterated_matrix {
		for ind, p := range v {
			sums[ind] += p
		}
	}
	fmt.Println(sums)

	total := 0.0
	for _, s := range sums {
		total += s
	}
	fmt.Println(total)

	keys := make([]int, 0, len(sums))

	for key := range sums {
		keys = append(keys, key)
	}

	fmt.Println(sums)
	fmt.Println(keys)

	sort.SliceStable(keys, func(i, j int) bool {
		return sums[keys[i]] > sums[keys[j]]
	})

	fmt.Println(keys)

}
