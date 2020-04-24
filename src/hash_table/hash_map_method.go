package hash_table

// 前置条件，只有一种结果
func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for index, val := range nums {
		if _, ok := numMap[val]; ok {
			numMap[val] += index
			continue
		}
		numMap[val] = index
	}
	for i := 0; i < len(nums); i++ {
		otherVal := target - nums[i]
		otherIdx, ok := numMap[otherVal]
		if !ok {
			continue
		}
		if nums[i] == otherVal {	// 这里考虑有重复值的情况
			otherIdx -= i
			if otherIdx == i {
				continue
			}
			if nums[i] != nums[otherIdx] {
				continue
			}
			return []int{i, otherIdx}
		}
		return []int{i, otherIdx}
	}
	return []int{-1, -1}
}