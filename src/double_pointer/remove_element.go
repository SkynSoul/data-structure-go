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

func RemoveDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    dstLen, val := 0, nums[0] - 1
    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
            nums[dstLen] = nums[i]
            val = nums[i]
            dstLen++
        }
    }
    return dstLen
}