package double_pointer

// 数组两两成对，每对取较小值，找出最大相加之和
func ArrayPairSum(nums []int) int {
	quickSort(nums)
	sum := 0
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}

func quickSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	i, j := 0, len(nums) - 1
	base, dir := nums[i], 0
	for i < j {
		if dir == 0 {
			if nums[j] <= base {
				nums[i] = nums[j]
				i++
				dir = 1
			} else {
				j--
			}
		} else {
			if nums[i] > base {
				nums[j] = nums[i]
				j--
				dir = 0
			} else {
				i++
			}
		}
	}
	nums[i] = base
	quickSort(nums[0:i])
	quickSort(nums[i+1:])
}

func quickSortWithStack(nums []int) {
	if len(nums) < 2 {
		return
	}
	stackHelper := make([][]int, 0, len(nums))
	stackHelper = append(stackHelper, nums[:])
	for len(stackHelper) > 0 {
		curArr := stackHelper[len(stackHelper) - 1]
		stackHelper = stackHelper[:len(stackHelper) - 1]
		if len(curArr) < 2 {
			continue
		}
		i, j, base := 0, len(curArr) - 1, curArr[0]
		for i < j {
			for i < j && curArr[j] > base {
				j--
			}
			if i < j {
				curArr[i] = curArr[j]
			}
			for i < j && curArr[i] <= base {
				i++
			}
			if i < j {
				curArr[j] = curArr[i]
			}
		}
		curArr[i] = base
		stackHelper = append(stackHelper, curArr[:i])
		stackHelper = append(stackHelper, curArr[i + 1:])
	}
}