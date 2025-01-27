package sort

import "testing"

func TestInsertionSortAsc(t *testing.T) {
	test := []int{1, 2, 3, 4, 5}

	r := insertionSort([]int{1, 5, 3, 2, 4}, true)

	if len(r) != len(test) {
		t.Fail()
	}

	for i, e := range test {
		if e != r[i] {
			t.Fail()
		}
	}
}

func TestInsertionSortDesc(t *testing.T) {
	test := []int{100, 90, 85, 60, 32, 18, 9, 0}

	r := insertionSort([]int{60, 85, 90, 18, 0, 100, 9, 32}, false)

	if len(r) != len(test) {
		t.Fail()
	}

	for i, e := range test {
		if e != r[i] {
			t.Fail()
		}
	}
}
