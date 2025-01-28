package binaries

import (
	"testing"
)

func TestBinaryIntSumCase1(t *testing.T) {
	r := BinaryIntSum([]bool{true}, []bool{true})
	expected := []bool{true, false}
	for i := 0; i < 2; i++ {
		if r[i] != expected[i] {
			t.Fail()
		}
	}
}

func TestBinaryIntSumCase2(t *testing.T) {
	r := BinaryIntSum([]bool{true, true, false, true}, []bool{true, false, true, true})
	expected := []bool{true, true, false, false, false}
	for i := 0; i < 2; i++ {
		if r[i] != expected[i] {
			t.Fail()
		}
	}
}

func TestBinaryIntSum(t *testing.T) {
	a := []bool{true, false, true, false, false, true, false, false, true, false, false, true, true, true, true, false, true, true, false, false, true, true, false, true, true, true}
	b := []bool{true, false, true, true, true, false, false, false, true, true, false, true, true, true, false, false, false, true, false, false, true, true, false, true, true, true}

	r := BinaryIntSum(a, b)

	expected := []bool{true, false, true, false, true, true, true, false, true, false, true, true, true, true, false, true, true, false, false, false, true, true, false, true, true, true, false}

	for i := 0; i < 2; i++ {
		if r[i] != expected[i] {
			t.Errorf("Test failed at position %d: expected %v, got %v", i, expected[i], r[i])
		}
	}
}

func TestBitArrayToInt64TestCase1(t *testing.T) {

	// 10100100100111101100110111
	if BitArrayToInt64([]bool{true, false, true, false, false, true, false, false, true, false, false, true, true, true, true, false, true, true, false, false, true, true, false, true, true, true}) != 43154231 {
		t.Fail()
	}
}

func TestInt64ToBitArray(t *testing.T) {
	r := Int64ToBitArray(32)[56:64] // only last 8 is enough for testing
	expected := []bool{false, false, true, false, false, false, false, false}

	for i := 0; i < len(expected); i++ {
		if expected[i] != r[i] {
			t.Fail()
		}
	}
}
