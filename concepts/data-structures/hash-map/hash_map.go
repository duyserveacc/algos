// Package hashmap demonstrates lookup-based uses of Go maps.
package hashmap

// PairSumIndices returns indices of two distinct values whose sum is target.
// The boolean is false when no pair exists.
func PairSumIndices(values []int, target int) (int, int, bool) {
	// Map every previously visited value to one of its earlier indices.
	indexByValue := make(map[int]int)
	// Visit each value once from left to right.
	for index, value := range values {
		// Calculate the value needed to complete the target.
		complement := target - value
		// Look up the complement and whether it exists as two explicit results.
		earlierIndex, found := indexByValue[complement]
		// Continue into this branch only when an earlier complement was found.
		if found {
			// Return both distinct indices and report that a pair was found.
			return earlierIndex, index, true
		}
		// Insert after lookup so the current element cannot match itself.
		indexByValue[value] = index
	}

	// Return the zero indices with false when the input contains no valid pair.
	return 0, 0, false
}
