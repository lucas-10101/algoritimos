package sort

// 1,2,3,4,6,5
// 1,2,3,4,6,6
// 0,1,2,3,5,6
func InsertionSort(v []int) []int {

	for i := 1; i < len(v); i++ {
		k := v[i]
		j := i - 1

		for j >= 0 && v[j] > k {
			v[j+1] = v[j]
			j--
		}

		v[j+1] = k
	}

	return v
}
