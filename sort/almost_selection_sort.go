package sort

import (
	"strconv"
)

func sort(v []int) []int {
	for i := 0; i < len(v); i += 1 {
		for j := i + 1; j < len(v); j += 1 {
			if v[i] > v[j] {
				v[i], v[j] = v[j], v[i]
			}
		}
	}

	return v
}

func sliceToString(v []int) string {

	s := ""

	for _, e := range v {
		s += strconv.FormatInt(int64(e), 10)
	}

	return s
}
