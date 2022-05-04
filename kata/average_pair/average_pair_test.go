package average_pair

import "testing"

func TestAveragePairReturnsFalseForArrayWithLessThanTwoElements(t *testing.T) {
	for _, tc := range makeTestCasesWithInputArrayOfLessThanTwoElementsInSize() {
		exp := tc.exp
		inputArr := tc.inputArr
		inputAvg := tc.inputAvg
		got := averagePair(inputArr, inputAvg)
		if exp != got {
			t.Errorf(
				"Expected %v, got %v instead for %v with average pair value %v",
				exp,
				got,
				inputArr,
				inputAvg,
			)
		}
	}
}

func TestAveragePairReturnsTrueForArrayWithAveragePair(t *testing.T) {
	for _, tc := range makeTestCasesWithAveragePair() {
		exp := tc.exp
		inputArr := tc.inputArr
		inputAvg := tc.inputAvg
		got := averagePair(inputArr, inputAvg)
		if exp != got {
			t.Errorf(
				"Expected %v, got %v instead for %v with average pair value %v",
				exp,
				got,
				inputArr,
				inputAvg,
			)
		}
	}
}

func TestAveragePairReturnsFalseForArrayWithoutAveragePair(t *testing.T) {
	for _, tc := range makeTestCasesWithoutAveragePair() {
		exp := tc.exp
		inputArr := tc.inputArr
		inputAvg := tc.inputAvg
		got := averagePair(inputArr, inputAvg)
		if exp != got {
			t.Errorf(
				"Expected %v, got %v instead for %v with average pair value %v",
				exp,
				got,
				inputArr,
				inputAvg,
			)
		}
	}
}

// Helpers

type testCase struct {
	exp      bool
	inputArr []int
	inputAvg float32
}

func makeTestCasesWithInputArrayOfLessThanTwoElementsInSize() []testCase {
	return []testCase{
		{
			exp:      false,
			inputArr: []int{},
			inputAvg: 2.0,
		},
		{
			exp:      false,
			inputArr: []int{1},
			inputAvg: 2.1,
		},
	}
}

func makeTestCasesWithAveragePair() []testCase {
	return []testCase{
		{
			exp:      true,
			inputArr: []int{1, 2, 3},
			inputAvg: 2.5,
		},
		{
			exp:      true,
			inputArr: []int{1, 3, 3, 5, 6, 7, 10, 12, 19},
			inputAvg: 8.0,
		},
	}
}

func makeTestCasesWithoutAveragePair() []testCase {
	return []testCase{
		{
			exp:      false,
			inputArr: []int{1, 2, 3},
			inputAvg: 2.7,
		},
		{
			exp:      false,
			inputArr: []int{1, 3, 3, 5, 6, 7, 10, 12, 19},
			inputAvg: 17.4,
		},
		{
			exp:      false,
			inputArr: []int{-1, 0, 3, 4, 5, 6},
			inputAvg: 4.1,
		},
	}
}
