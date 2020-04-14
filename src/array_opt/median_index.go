package array_opt

func PivotIndex(nums []int) int {
	numLen := len(nums)
	leftArr := make([]int, numLen)
	rightArr := make([]int, numLen)
	sum := 0
	for i := 1; i < numLen; i++ {
		sum += nums[i - 1]
		leftArr[i] = sum
	}
	sum = 0
	for i := numLen - 2; i >= 0; i-- {
		sum += nums[i + 1]
		rightArr[i] = sum
	}
	for i := 0; i < numLen; i++ {
		if leftArr[i] == rightArr[i] {
			return i
		}
	}
	return -1
}

// sum = leftSum + nums[i] + rightSum
func PivotIndex2(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	leftSum := 0
	for i := 0; i < len(nums); i++ {
		if leftSum == sum - leftSum - nums[i] {
			return i
		}
		leftSum += nums[i]
	}
	return -1
}