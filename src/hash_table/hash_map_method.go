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

func IsIsomorphic(s string, t string) bool {
	ruleArr1 := getStringRule(s)
	ruleArr2 := getStringRule(t)
	if len(ruleArr1) != len(ruleArr2) {
		return false
	}
	for i := 0; i < len(ruleArr1); i++ {
		if len(ruleArr1[i]) != len(ruleArr2[i]) {
			return false
		}
		for j := 0; j < len(ruleArr1[i]); j++ {
			if ruleArr1[i][j] != ruleArr2[i][j] {
				return false
			}
		}
	}
	return true
}

func getStringRule(s string) [][]int {
	charMap := make(map[byte]int)
	ruleArr := make([][]int, 0, len(s))
	charIdx := 0
	for i := 0; i < len(s); i++ {
		tarIdx, ok := charMap[s[i]]
		if !ok {
			ruleArr = append(ruleArr, []int{i})
			charMap[s[i]] = charIdx
			charIdx++
			continue
		}
		ruleArr[tarIdx] = append(ruleArr[tarIdx], i)
	}
	return ruleArr
}

func IsIsomorphic2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	return doIsomorphic(s, t) && doIsomorphic(t, s)
}

func doIsomorphic(s string, t string) bool {
	charMap := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		reChar, ok := charMap[t[i]]
		if !ok {
			charMap[t[i]] = s[i]
			continue
		}
		if s[i] != reChar {
			return false
		}
	}
	return true
}

func FindRestaurant(list1 []string, list2 []string) []string {
	nameMap := make(map[string]int)
	for index, name := range list1 {
		nameMap[name] = index
	}
	retMap := make(map[int][]string)
	minSum := -1
	for index, name := range list2 {
		idx, ok := nameMap[name]
		if !ok {
			continue
		}
		idxSum := index + idx
		if minSum < 0 || idxSum < minSum {
			minSum = idxSum
		}
		tarArr, ok := retMap[idxSum]
		if !ok {
			retMap[idxSum] = []string{name}
			continue
		}
		retMap[idxSum] = append(tarArr, name)
	}
	if minSum < 0 {
		return []string{}
	}
	return retMap[minSum]
}