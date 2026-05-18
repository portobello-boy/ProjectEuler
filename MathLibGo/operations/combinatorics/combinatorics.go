package combinatorics

func Permute[V any](items []V, k int) [][]V {
	if k == 0 {
		return [][]V{{}}
	}
	if len(items) < k {
		return nil
	}
	result := make([][]V, 0)
	for i := range items {
		rest := make([]V, 0)
		rest = append(rest, items[:i]...)
		rest = append(rest, items[i+1:]...)
		subPermutations := Permute(rest, k-1)
		for _, sub := range subPermutations {
			result = append(result, append([]V{items[i]}, sub...))
		}
	}
	return result
}
