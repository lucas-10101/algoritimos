package search

import "testing"

func TestLinearSearchTestCase1(t *testing.T) {
	if linearSearch([]int{1, 2, 3, 4, 5}, 3) != 2 {
		t.Fail()
	}
}
func TestLinearSearchTestCase2(t *testing.T) {
	if linearSearch([]int{1, 2, 3, 4, 5}, 10) != -1 {
		t.Fail()
	}
}
func TestLinearSearchTestCase3(t *testing.T) {
	if linearSearch([]int{1, 2, 3, 4, 5}, 5) != 4 {
		t.Fail()
	}
}
