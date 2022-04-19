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

	lookup := make(map[rune]int)

	for _, value := range s1 {
		lookup[value]++
	}

	for _, value := range s2 {
		var count int
		var ok bool
		if count, ok = lookup[value]; !ok {
			return false
		}
		if count == 0 {
			return false
		}
		lookup[value]--
	}
	return true
}
