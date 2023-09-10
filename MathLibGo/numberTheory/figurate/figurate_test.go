package figurate

import "testing"

func TestTriangular(t *testing.T) {
	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedTriangular := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}

	for index, input := range inputs {
		result, _ := Triangular(input)

		if result != expectedTriangular[index] {
			t.Errorf("Got: %d, wanted: %d", result, expectedTriangular[index])
		}
	}
}

func TestTriangularInvalid(t *testing.T) {
	result, err := Triangular(-1)

	if err != nil && result != -1 {
		t.Errorf("Expected error but didn't get one!")
	}
}

func TestPolygonal(t *testing.T) {
	expectedPolygonals := map[int][]int{
		3: {1, 3, 6, 10, 15, 21, 28, 36, 45, 55},     // Triangular
		4: {1, 4, 9, 16, 25, 36, 49, 64, 81, 100},    // Square
		5: {1, 5, 12, 22, 35, 51, 70, 92, 117, 145},  // Pentagonal
		6: {1, 6, 15, 28, 45, 66, 91, 120, 153, 190}, // Hexagonal
	}

	for r, expectedOutputs := range expectedPolygonals {
		t.Logf("r: %d, expectedOutputs: %v", r, expectedOutputs)
		for index, expectedOutput := range expectedOutputs {
			result, _ := Polygonal(r, index+1)

			if result != expectedOutput {
				t.Errorf("Got: %d, wanted: %d, input was: %d, %d", result, expectedOutput, r, index+1)
			}
		}
	}
}
