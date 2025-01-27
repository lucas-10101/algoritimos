package sort

func insertionSort(v []int, asc bool) []int {

	for i := 1; i < len(v); i++ {
		k := v[i]
		j := i - 1

		for j >= 0 && (asc && v[j] > k || !asc && v[j] < k) {
			v[j+1] = v[j]
			j--
		}

		v[j+1] = k
	}

	return v
}
