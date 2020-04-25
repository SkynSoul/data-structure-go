package hash_table

import (
	"sort"
)

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

func FirstUniqChar(s string) int {
	charMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := charMap[s[i]]; ok {
			charMap[s[i]] = -1
			continue
		}
		charMap[s[i]] = i
	}
	minIdx := -1
	for _, val := range charMap {
		if val == -1 {
			continue
		}
		if minIdx < 0 || val < minIdx {
			minIdx = val
		}
	}
	return minIdx
}

func Intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	charMap := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		if _, ok := charMap[nums1[i]]; ok {
			charMap[nums1[i]] += 1
			continue
		}
		charMap[nums1[i]] = 1
	}
	ret := make([]int, 0)
	for i := 0; i < len(nums2) && len(charMap) > 0; i++ {
		if count, ok := charMap[nums2[i]]; ok {
			ret = append(ret, nums2[i])
			charMap[nums2[i]] -= 1
			if count == 1 {
				delete(charMap, nums2[i])
			}
		}
	}
	return ret
}

func ContainsNearbyDuplicate(nums []int, k int) bool {
	numMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if lstIdx, ok := numMap[nums[i]]; ok {
			if i - lstIdx <= k {
				return true
			}
		}
		numMap[nums[i]] = i
	}
	return false
}

func GroupAnagrams(strs []string) [][]string {
	strMap := make(map[string][]string)
	for _, str := range strs {
		strArr := []byte(str)
		sort.Slice(strArr, func(i, j int) bool {
			return strArr[i] > strArr[j]
		})
		if tarArr, ok := strMap[string(strArr)]; ok {
			strMap[string(strArr)] = append(tarArr, str)
			continue
		}
		strMap[string(strArr)] = []string{str}
	}
	ret := make([][]string, 0, len(strMap))
	for _, val := range strMap {
		ret = append(ret, val)
	}
	return ret
}