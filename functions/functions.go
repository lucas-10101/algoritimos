package functions

import "strconv"

func IntegerSliceToString(v []int) string {

	s := ""

	for _, e := range v {
		s += strconv.FormatInt(int64(e), 10)
	}

	return s
}
