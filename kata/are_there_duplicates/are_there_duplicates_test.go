package are_there_duplicates

import "testing"

func TestAreThereDuplicatesReturnsTrueForPositiveCase(t *testing.T) {
	got := areThereDuplicates(1, 2, 3, 4, 5, 6, 1, 7, 8, 9)
	if !got {
		t.Errorf("Expected TRUE got %v instaed", got)
	}
}

func TestAreThereDuplicatesReturnsFalseForNegativeCase(t *testing.T) {
	got := areThereDuplicates(1, 2, 3)
	if got {
		t.Errorf("Expected false got %v instaed", got)
	}
}
