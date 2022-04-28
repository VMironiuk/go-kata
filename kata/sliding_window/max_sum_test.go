package sliding_window

import "testing"

func TestMaxSumReturnsZeroForEmptyArray(t *testing.T) {
	exp := 0
	got := maxSum([]int{}, 3)
	if exp != got {
		t.Errorf("Expected %v, got %v instead", exp, got)
	}
}

func TestMaxSumReturnsZeroIfCountToSumGreaterThenArrayLength(t *testing.T) {
	exp := 0
	got := maxSum([]int{1, 2}, 3)
	if exp != got {
		t.Errorf("Expected %v, got %v instead", exp, got)
	}
}

func TestMaxSumReturnsValidSum(t *testing.T) {
	exp := 19
	got := maxSum([]int{2, 6, 9, 2, 1, 8, 5, 6, 3}, 3)
	if exp != got {
		t.Errorf("Expected %v, got %v instead", exp, got)
	}
}
