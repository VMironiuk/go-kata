package multiple_pointers

func countUniqueValues(arr []int) int {
	return countUniqueValuesWithoutCounter(arr)
}

func countUniqueValuesWithCounter(arr []int) int {
	if len(arr) <= 1 {
		return len(arr)
	}

	counter := 1
	for i, j := 0, 1; j < len(arr); i, j = i+1, j+1 {
		if arr[i] != arr[j] {
			counter++
		}
	}
	return counter
}

func countUniqueValuesWithoutCounter(arr []int) int {
	if len(arr) <= 1 {
		return len(arr)
	}

	i := 0
	for j := i + 1; j < len(arr); j++ {
		if arr[i] != arr[j] {
			i++
			arr[i] = arr[j]
		}
	}

	return i + 1
}
