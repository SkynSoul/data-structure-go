package arr_string_opt

// 时间复杂度O(N * k)
func Rotate(nums []int, k int) {
    numsLen := len(nums)
    k = k % numsLen
    for i := 0; i < k; i++ {
        for j := 0; j < numsLen - 1; j++ {
            nums[j], nums[numsLen - 1] = nums[numsLen - 1], nums[j]
        }
    }
}

// 1、翻转所有元素
// 2、翻转前k个元素
// 3、翻转k之后的元素
// 时间复杂度O(N)
func Rotate2(nums []int, k int) {
    reverse(nums)
    reverse(nums[:k])
    reverse(nums[k:])
}

func reverse(nums []int) {
    left, right := 0, len(nums) - 1
    for left < right {
        nums[left], nums[right] = nums[right], nums[left]
        left++
        right--
    }
}