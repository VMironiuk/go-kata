package sliding_window

func maxSum(arr []int, count int) int {
	if len(arr) < count {
		return 0
	}

	maxSum := 0
	tmpSum := 0

	for i := 0; i < count; i++ {
		maxSum += arr[i]
	}

	tmpSum = maxSum

	for i := count; i < len(arr); i++ {
		tmpSum = tmpSum - arr[i-count] + arr[i]
		if maxSum < tmpSum {
			maxSum = tmpSum
		}
	}

	return maxSum
}
