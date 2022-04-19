package frequency_counter

func same(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	am := make(map[int]int)
	bm := make(map[int]int)
	for _, k := range a {
		am[k]++
	}
	for _, k := range b {
		bm[k]++
	}
	for k, v := range am {
		var bv int
		var ok bool
		if bv, ok = bm[k*k]; !ok {
			return false
		}
		if bv != v {
			return false
		}
	}
	return true
}

func isValidAnagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	var freqCounts1 = make(map[rune]int)
	var freqCounts2 = make(map[rune]int)
	for _, val := range s1 {
		freqCounts1[val]++
	}
	for _, val := range s2 {
		freqCounts2[val]++
	}
	for key, val := range freqCounts1 {
		var freqCounts2Val int
		var ok bool
		if freqCounts2Val, ok = freqCounts2[key]; !ok {
			return false
		}
		if val != freqCounts2Val {
			return false
		}
	}
	return true
}
