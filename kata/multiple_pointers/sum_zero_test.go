package multiple_pointers

import "testing"

func TestSumZero(t *testing.T) {
	testData := makeSumZeroTestData()
	for _, data := range testData {
		arr := data.arr
		exp := data.exp
		got := sumZero(arr)
		if !isSlicesEqual(exp, got) {
			t.Errorf("Expected %v, got %v instead", exp, got)
		}
	}
}

// Helpers

type sumZeroTestData struct {
	arr []int
	exp []int
}

func makeSumZeroTestData() []sumZeroTestData {
	return []sumZeroTestData{
		{
			[]int{-4, -3, -2, -1, 0, 1, 2, 3, 5},
			[]int{-3, 3},
		},
		{
			[]int{-4, -2, -1, 0, 1, 3, 4, 5, 6},
			[]int{-4, 4},
		},
		{
			[]int{-4, -3, -2, -1, 0, 6, 7, 9, 11},
			[]int{},
		},
		{
			[]int{-4, -2, -1, 0},
			[]int{},
		},
	}
}

func isSlicesEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
