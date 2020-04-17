package double_pointer

func TwoSum(numbers []int, target int) []int {
    for i := 0; i < len(numbers); i++ {
        if j := binarySearch(numbers[i:], target - numbers[i]); j != -1 {
            return []int{i + 1, i + j + 1}
        }
    }
    return []int{-1, -1}
}

func binarySearch(nums []int, target int) int {
    left, right := 0, len(nums)
    for left < right {
        median := (left + right) / 2
        if nums[median] == target {
            return median
        }
        if nums[median] < target {
            left = median + 1
            continue
        }
        right = median
    }
    return -1
}

func TwoSum2(numbers []int, target int) []int {
    left, right := 0, len(numbers) - 1
    for left < right {
        if numbers[left] + numbers[right] == target {
            return []int{left + 1, right + 1}
        }
        if numbers[left] + numbers[right] < target {
            left++
        } else {
            right--
        }
    }
    return []int{-1, -1}
}