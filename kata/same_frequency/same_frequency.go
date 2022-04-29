package same_frequency

func sameFrequency(lhs, rhs int) bool {
	var lhsArr []int
	var rhsArr []int
	for lhs > 0 || rhs > 0 {
		if lhs > 0 {
			lhsArr = append(lhsArr, lhs%10)
			lhs /= 10
		}
		if rhs > 0 {
			rhsArr = append(rhsArr, rhs%10)
			rhs /= 10
		}
	}

	if len(lhsArr) != len(rhsArr) {
		return false
	}

	lhsMap := make(map[int]int)
	rhsMap := make(map[int]int)
	for i := 0; i < len(lhsArr); i++ {
		lhsMap[lhsArr[i]]++
		rhsMap[rhsArr[i]]++
	}

	for key, value := range lhsMap {
		val, ok := rhsMap[key]
		if !ok {
			return false
		}
		if val != value {
			return false
		}
	}

	return true
}
