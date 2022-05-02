package are_there_duplicates

func areThereDuplicates(entities ...int) bool {
	set := make(map[int]struct{})
	for _, ent := range entities {
		if _, ok := set[ent]; ok {
			return true
		}
		set[ent] = struct{}{}
	}
	return false
}
