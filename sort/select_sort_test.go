package sort

import "testing"

func TestSelectSortCase1(t *testing.T) {

	actual := selectSort([]int{1, 3, 5, 7, 9})
	expected := []int{1, 3, 5, 7, 9}

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Errorf("Expected: %d but: %d was given (%s)", expected[i], actual[i], t.Name())
		}
	}
}

func TestSelectSortCase2(t *testing.T) {

	actual := selectSort([]int{1})
	expected := []int{1}

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Errorf("Expected: %d but: %d was given (%s)", expected[i], actual[i], t.Name())
		}
	}
}

func TestSelectSortCase3(t *testing.T) {

	actual := selectSort([]int{9, 8, 7, 6, 5, 4, 3, 2, 1})
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Errorf("Expected: %d but: %d was given (%s)", expected[i], actual[i], t.Name())
		}
	}
}

func TestSelectSortCase4(t *testing.T) {

	actual := selectSort([]int{1, 3, 2, 4, 6, 5, 7, 9, 8})
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Errorf("Expected: %d but: %d was given (%s)", expected[i], actual[i], t.Name())
		}
	}
}

func TestSelectSortCase5(t *testing.T) {

	actual := selectSort([]int{5, 5, 4, 4, 3, 3, 2, 2, 1, 1})
	expected := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Errorf("Expected: %d but: %d was given (%s)", expected[i], actual[i], t.Name())
		}
	}
}
