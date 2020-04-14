package array_opt

func DominantIndex(nums []int) int {
	maxIdx, maxDouble := 0, 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[maxIdx] {
			maxDouble = 2 * nums[maxIdx]
			maxIdx = i
		} else if maxDouble < 2 * nums[i] {
			maxDouble = 2 * nums[i]
		}
	}
	if maxDouble <= nums[maxIdx] {
		return maxIdx
	}
	return -1
}