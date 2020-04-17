package double_pointer

func RemoveElement(nums []int, val int) int {
    dstLen := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
            nums[dstLen] = nums[i]
            dstLen++
        }
    }
    return dstLen
}