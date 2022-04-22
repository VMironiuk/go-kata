package multiple_pointers

import "testing"

func TestCountUniqueValuesReturnsZeroForEmptyArray(t *testing.T) {
	testCases := makeCountUniqueValuesTestData()
	for _, testCase := range testCases {
		input := testCase.input
		exp := testCase.exp
		got := countUniqueValues(input)
		if got != exp {
			t.Errorf("Expected %v, got %v instead for input: %v", exp, got, input)
		}
	}
}

// Helpers

type countUniqueValuesTestCase struct {
	input []int
	exp   int
}

func makeCountUniqueValuesTestData() []countUniqueValuesTestCase {
	return []countUniqueValuesTestCase{
		{
			[]int{},
			0,
		},
		{
			[]int{1},
			1,
		},
		{
			[]int{1, 1},
			1,
		},
		{
			[]int{1, 2},
			2,
		},
		{
			[]int{1, 1, 1, 1, 1, 2},
			2,
		},
		{
			[]int{1, 2, 3, 4, 4, 4, 7, 7, 12, 12, 13},
			7,
		},
		{
			[]int{-2, -1, -1, 0, 1},
			4,
		},
	}
}
