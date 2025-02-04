package sort

func selectSort(v []int) []int {

	for i := 0; i < len(v); i++ {
		actual, newPos := v[i], i

		for j := i; j < len(v); j++ {
			if v[newPos] > v[j] {
				newPos = j
			}
		}

		v[i] = v[newPos]
		v[newPos] = actual
	}

	return v
}
