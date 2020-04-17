package double_pointer

func FindMaxConsecutiveOnes(nums []int) int {
    ret := 0
    for i := 0; i < len(nums); {
        if nums[i] == 1 {
            length, j := 1, i + 1
            for j < len(nums) {
                if nums[j] != 1 {
                    j++
                    break
                }
                length++
                j++
            }
            i = j
            if ret < length {
                ret = length
            }
            continue
        }
        i++
    }
    return ret
}