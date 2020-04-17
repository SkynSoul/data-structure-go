package arr_string_opt

func Rotate(nums []int, k int) {
    numsLen := len(nums)
    k = k % numsLen
    for i := 0; i < k; i++ {
        for j := 0; j < numsLen - 1; j++ {
            nums[j], nums[numsLen - 1] = nums[numsLen - 1], nums[j]
        }
    }
}