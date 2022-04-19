package frequency_counter

import "testing"

func TestSame(t *testing.T) {
	if same([]int{1, 3, 2, 2}, []int{9, 4, 1, 4}) != true {
		t.Error("Expected true")
	}
}

func TestIsValidAnagramReturnsTrueForBothEmptyStrings(t *testing.T) {
	if isValidAnagram("", "") != true {
		t.Error("Expected true for both empty strings")
	}
}

func TestIsValidAnagramReturnsTrueForStringsWithSameSetOfCharacters(t *testing.T) {
	anagrams := makeAnagrams()
	for _, pair := range anagrams {
		if isValidAnagram(pair.s1, pair.s2) != true {
			t.Errorf("Expected true for %s and %s", pair.s1, pair.s2)
		}
	}
}

func TestIsValidAnagramReturnsFalseForStringsWithDifferentSetOfCharacters(t *testing.T) {
	anagrams := makeInvalidAnagrams()
	for _, pair := range anagrams {
		if isValidAnagram(pair.s1, pair.s2) != false {
			t.Errorf("Expected false for %s and %s", pair.s1, pair.s2)
		}
	}
}

// Helpers

type testSet struct {
	s1 string
	s2 string
}

func makeAnagrams() []testSet {
	return []testSet{
		{
			"anagram",
			"nagaram",
		},
		{
			"qwerty",
			"qeywrt",
		},
		{
			"texttwisttime",
			"timetwisttext",
		},
	}
}

func makeInvalidAnagrams() []testSet {
	return []testSet{
		{
			"aaz",
			"zza",
		},
		{
			"awesome",
			"awesom",
		},
		{
			"texttwisttimepre",
			"texttwisttinerpe",
		},
	}
}
