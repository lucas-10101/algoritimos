package main

import (
	"fmt"

	"github.com/lucas-10101/algoritimos/functions"
)

func sorting(s []int) []int {

	for i := 0; i < len(s); i++ {
		v := s[i]
		j := i - 1

		for j >= 0 && s[j] > v {
			s[j+1] = s[j]
			j--
		}

		s[j+1] = v
	}

	return s

}

func main() {

	fmt.Println(functions.IntegerSliceToString(sorting([]int{2, 4, 5, 1, 6, 3, 0, 8, 7, 9})))
}
