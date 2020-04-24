package hash_table

import "math"

func ContainsDuplicate(nums []int) bool {
    mapHelper := make(map[int]bool)
    for i := 0; i < len(nums); i++ {
        if _, ok := mapHelper[nums[i]]; ok {
            return true
        }
        mapHelper[nums[i]] = true
    }
    return false
}

// 前提条件：重复数字的个数为偶数
func SingleNumber(nums []int) int {
    mapHelper := make(map[int]bool)
    for i := 0; i < len(nums); i++ {
        if _, ok := mapHelper[nums[i]]; ok {
            delete(mapHelper, nums[i])
            continue
        }
        mapHelper[nums[i]] = true
    }
    for key := range mapHelper {
        return key
    }
    return -1
}

func SingleNumberXOR(nums []int) int {
    ret := 0
    for _, val := range nums {
        ret ^= val
    }
    return ret
}

func Intersection(nums1 []int, nums2 []int) []int {
    ret := make([]int, 0)
    mapHelper := make(map[int]int)
    for _, val := range nums1 {
        if _, ok := mapHelper[val]; !ok {
            mapHelper[val] = 1
        }
    }
    for _, val := range nums2 {
        if _, ok := mapHelper[val]; ok {
            mapHelper[val] += 1
        }
    }
    for key, val := range mapHelper {
        if val > 1 {
            ret = append(ret, key)
        }
    }
    return ret
}

func IsHappy(n int) bool {
    numArr := make([]int, 0)
    numMap := make(map[int]bool)
    for {
        numArr = numArr[:0]
        for n > 0 {
            numArr = append(numArr, n % 10)
            n /= 10
        }
        sum := 0
        for _, val := range numArr {
            temp := val * val
            // 这里其实没有必要，全为9的13位数字，最后结果也只是1053
            if math.MaxInt64 - sum < temp {
                return false
            }
            sum += val * val
        }
        if sum == 1 {
            return true
        }
        if _, ok := numMap[sum]; ok {
            return false
        }
        numMap[sum] = true
        n = sum
    }
}