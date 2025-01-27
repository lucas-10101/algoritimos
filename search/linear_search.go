package search

// Try find index of target t in slice s, or return -1 if not found
func linearSearch(s []int, t int) int {
	for i, e := range s {
		if e == t {
			return i
		}
	}
	return -1
}
