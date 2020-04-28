package hash_table

import (
	"math"
	"sort"
	"strconv"
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

func IsValidSudoku(board [][]byte) bool {
	rowMap := make(map[byte]map[byte]bool)
	colMap := make(map[byte]map[byte]bool)
	cellMap := make(map[byte]map[byte]bool)
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[x]); y++ {
			if board[x][y] == '.' {
				continue
			}
			rowNums, ok := rowMap[byte(x)]
			if !ok {
				rowMap[byte(x)] = map[byte]bool{board[x][y]: true}
			} else if _, ok := rowNums[board[x][y]]; ok {
				return false
			} else {
				rowNums[board[x][y]] = true
			}
			colNums, ok := colMap[byte(y)]
			if !ok {
				colMap[byte(y)] = map[byte]bool{board[x][y]: true}
			} else if _, ok := colNums[board[x][y]]; ok {
				return false
			} else {
				colNums[board[x][y]] = true
			}
			cellKey := ((x / 3) * 3) + (y / 3)
			cellNums, ok := cellMap[byte(cellKey)]
			if !ok {
				cellMap[byte(cellKey)] = map[byte]bool{board[x][y]: true}
			} else if _, ok := cellNums[board[x][y]]; ok {
				return false
			} else {
				cellNums[board[x][y]] = true
			}
		}
	}
	return true
}

func IsValidSudokuByArr(board [][]byte) bool {
	rowArr := make([][]byte, 9)
	colArr := make([][]byte, 9)
	cellArr := make([][]byte, 9)
	for i := 0; i < 9; i++ {
		rowArr[i] = make([]byte, 10)
		colArr[i] = make([]byte, 10)
		cellArr[i] = make([]byte, 10)
	}
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[x]); y++ {
			if board[x][y] == '.' {
				continue
			}
			idx := board[x][y] - '0'
			rowNums := rowArr[x]
			if rowNums[idx] == 1 {
				return false
			}
			rowNums[idx] = 1

			colNums := colArr[y]
			if colNums[idx] == 1 {
				return false
			}
			colNums[idx] = 1
			
			cellKey := ((x / 3) * 3) + (y / 3)
			cellNums := cellArr[cellKey]
			if cellNums[idx] == 1 {
				return false
			}
			cellNums[idx] = 1
		}
	}
	return true
}

type TreeNode struct {
	Val     int
	Left    *TreeNode
	Right   *TreeNode
}

func FindDuplicateSubtrees(root *TreeNode) []*TreeNode {
	retMap := make(map[string][]*TreeNode)
	preorder(root, retMap)
	ret := make([]*TreeNode, 0)
	for _, tarArr := range retMap {
		if len(tarArr) > 1 {
			ret = append(ret, tarArr...)
		}
	}
	return ret
}

func preorder(root *TreeNode, retMap map[string][]*TreeNode) string {
	if root == nil {
		return ""
	}
	curStr := strconv.Itoa(root.Val)
	leftStr := preorder(root.Left, retMap)
	rightStr := preorder(root.Right, retMap)
	curStr = curStr + "[L" + leftStr + "]"
	curStr = curStr + "[R" + rightStr + "]"
	nodeArr, ok := retMap[curStr]
	if !ok {
		retMap[curStr] = []*TreeNode{root}
	} else {
		retMap[curStr] = append(nodeArr, root)
	}
	return curStr
}

func NumJewelsInStones(J string, S string) int {
	gemMap := make(map[byte]bool)
	for i := 0; i < len(J); i++ {
		gemMap[J[i]] = true
	}
	ok := false
	num := 0
	for i := 0; i < len(S); i++ {
		if _, ok = gemMap[S[i]]; ok {
			num++
		}
	}
	return num
}

func NumJewelsInStonesByArr(J string, S string) int {
	gemArr := make([]byte, 58)
	for i := 0; i < len(J); i++ {
		gemArr[J[i] - 'A'] = 1
	}
	num := 0
	for i := 0; i < len(S); i++ {
		if gemArr[S[i] - 'A'] == 1 {
			num++
		}
	}
	return num
}

func LengthOfLongestSubstring(s string) int {
	ret, preIdx := 0, 0
	sMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if idx, ok := sMap[s[i]]; ok {
			if ret < len(sMap) {
				ret = len(sMap)
			}
			for j := preIdx; j < idx; j++ {
				delete(sMap, s[j])
			}
			preIdx = idx + 1
		}
		sMap[s[i]] = i
	}
	if ret < len(sMap) {
		ret = len(sMap)
	}
	return ret
}

func LengthOfLongestSubstring2(s string) int {
	ret, preIdx := 0, 0
	sMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if idx, ok := sMap[s[i]]; ok && idx >= preIdx {
			ret = int(math.Max(float64(ret), float64(i - preIdx)))
			preIdx = idx + 1
		}
		sMap[s[i]] = i
	}
	return int(math.Max(float64(ret), float64(len(s) - preIdx)))
}

func LengthOfLongestSubstring3(s string) int {
	ret, preIdx := 0, 1
	sArr := [256]int{}
	for i := 0; i < len(s); i++ {
		if sArr[s[i]] >= preIdx {
			ret = int(math.Max(float64(ret), float64(i - preIdx + 1)))
			preIdx = sArr[s[i]] + 1
		}
		sArr[s[i]] = i + 1
	}
	return int(math.Max(float64(ret), float64(len(s) - preIdx + 1)))
}

func LengthOfLongestSubstring4(s string) int {
	ret, preIdx := 0, 0
	sMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if idx, ok := sMap[s[i]]; ok {
			ret = int(math.Max(float64(ret), float64(len(sMap))))
			for j := preIdx; j < idx; j++ {
				delete(sMap, s[j])
			}
			preIdx = idx + 1
		}
		sMap[s[i]] = i
	}
	return int(math.Max(float64(ret), float64(len(sMap))))
}