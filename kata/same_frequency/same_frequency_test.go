package same_frequency

import "testing"

func TestSameFrequencyPositiveCase(t *testing.T) {
	testCases := makePositiveTestCases()
	for _, testCase := range testCases {
		lhs := testCase.lhs
		rhs := testCase.rhs
		exp := testCase.exp
		got := sameFrequency(lhs, rhs)
		if exp != got {
			t.Errorf(
				"Expected %v, got %v instead for %v and %v numbers",
				exp,
				got,
				lhs,
				rhs,
			)
		}
	}
}

func TestSameFrequencyNegativeCase(t *testing.T) {
	testCases := makeNegativeTestCases()
	for _, testCase := range testCases {
		lhs := testCase.lhs
		rhs := testCase.rhs
		exp := testCase.exp
		got := sameFrequency(lhs, rhs)
		if exp != got {
			t.Errorf(
				"Expected %v, got %v instead for %v and %v numbers",
				exp,
				got,
				lhs,
				rhs,
			)
		}
	}
}

// Helpers

type testCase struct {
	lhs int
	rhs int
	exp bool
}

func makePositiveTestCases() []testCase {
	return []testCase{
		{
			182,
			281,
			true,
		},
		{
			3589578,
			5879385,
			true,
		},
	}
}

func makeNegativeTestCases() []testCase {
	return []testCase{
		{
			22,
			222,
			false,
		},
		{
			14,
			34,
			false,
		},
		{
			3589528,
			5879385,
			false,
		},
	}
}
