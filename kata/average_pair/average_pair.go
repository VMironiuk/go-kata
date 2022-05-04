package average_pair

func averagePair(inputArr []int, inputAvg float32) bool {
	if len(inputArr) < 2 {
		return false
	}

	i := 0
	j := len(inputArr) - 1
	for i < j {
		avg := (float32(inputArr[i]) + float32(inputArr[j])) / 2.0
		if avg == inputAvg {
			return true
		} else if avg < inputAvg {
			i++
		} else if avg > inputAvg {
			j--
		}
	}

	return false
}
