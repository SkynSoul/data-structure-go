package double_pointer

func MinSubArrayLen(s int, nums []int) int {
    minLen := len(nums) + 1
    for i := 0; i < len(nums); i++ {
        curLen, sum := 0, 0
        for j := i; j < len(nums); j++ {
            sum += nums[j]
            curLen++
            if sum >= s {
                break
            }
        }
        if sum >= s && curLen < minLen {
            minLen = curLen
        }
    }
    if minLen > len(nums) {
        return 0
    }
    return minLen
}

func MinSubArrayLen2(s int, nums []int) int {
    head, tail, sum, minLen := 0, 0, 0, len(nums) + 1
    for {
        if sum < s {
            if tail == len(nums) {
                break
            }
            sum += nums[tail]
            tail++
            continue
        }
        if sum >= s {
            if tail - head < minLen {
                minLen = tail - head
            }
            sum -= nums[head]
            head++
        }
    }
    if minLen > len(nums) {
        return 0
    }
    return minLen
}