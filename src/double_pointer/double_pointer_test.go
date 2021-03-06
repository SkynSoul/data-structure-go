package double_pointer

import (
	"sort"
	"testing"
)

func TestReverseString(t *testing.T) {
	arr := []byte{'h', 'e', 'l', 'l', 'o'}
	t.Logf("src str arr: %q\n", arr)
	ReverseString(arr)
	t.Logf("after reverse: %q\n", arr)
}

func TestArrayPairSum(t *testing.T) {
	nums := []int{1, 4, 3, 2}
	t.Logf("nums is %v, the ret is %v\n", nums, ArrayPairSum(nums))
}

func TestQuickSort(t *testing.T) {
	nums := []int{1, 4, 3, 2}
	t.Logf("nums is %v\n", nums)
	quickSort(nums)
	t.Logf("after sort, the nums is %v\n", nums)
	nums = []int{4, 7, 6, 5, 3, 2, 8, 1}
	t.Logf("nums is %v\n", nums)
	quickSort(nums)
	t.Logf("after sort, the nums is %v\n", nums)
	nums = []int{1, 4, 3, 2, 5, 7, 9, 2, 3, 6, 15, 24, 0, 4, 6, 20}
	t.Logf("nums is %v\n", nums)
	quickSort(nums)
	t.Logf("after sort, the nums is %v\n", nums)
	nums = []int{4, 7, 6, 5, 3, 2, 8, 1}
	t.Logf("nums is %v\n", nums)
	quickSortWithStack(nums)
	t.Logf("after sort, the nums is %v\n", nums)
	nums = []int{1, 4, 3, 2, 5, 7, 9, 2, 3, 6, 15, 24, 0, 4, 6, 20}
	t.Logf("nums is %v\n", nums)
	quickSortWithStack(nums)
	t.Logf("after sort, the nums is %v\n", nums)
}

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21}
	target := 0
	t.Logf("targe %d index is %d", target, binarySearch(nums, target))
	target = 1
	t.Logf("targe %d index is %d", target, binarySearch(nums, target))
	target = 15
	t.Logf("targe %d index is %d", target, binarySearch(nums, target))
	target = 11
	t.Logf("targe %d index is %d", target, binarySearch(nums, target))
	target = 20
	t.Logf("targe %d index is %d", target, binarySearch(nums, target))
	target = 100
	t.Logf("targe %d index is %d", target, binarySearch(nums, target))
}

func TestTwoSum(t *testing.T) {
	numbers := []int{2, 7, 11, 15}
	target := 9
	t.Logf("nums is %v, target is %d, ret is %v", numbers, target, TwoSum2(numbers, target))
	numbers = []int{5, 25, 75}
	target = 100
	t.Logf("nums is %v, target is %d, ret is %v", numbers, target, TwoSum2(numbers, target))
}

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	t.Logf("src nums is %v, remove element is %d, dst nums is %v", nums, 3, nums[:RemoveElement(nums, val)])
	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	val = 2
	t.Logf("src nums is %v, remove element is %d, dst nums is %v", nums, 3, nums[:RemoveElement(nums, val)])
}

func TestFindMaxConsecutiveOnes(t *testing.T) {
	nums := []int{1, 1, 0, 1, 1, 1}
	t.Logf("nums is %v, the max combo is %d", nums, FindMaxConsecutiveOnes(nums))
	nums = []int{1, 0, 1, 1, 0, 1}
	t.Logf("nums is %v, the max combo is %d", nums, FindMaxConsecutiveOnes(nums))
}

func TestMinSubArrayLen(t *testing.T) {
	nums := []int{2, 3, 1, 2, 4, 3, 8, 7, 1, 2, 5, 7, 3, 2, 1, 1, 2, 3, 1, 0, 0, 7}
	s := 7
	t.Logf("nums is %v, target sum is %d, the min array len is %d", nums, s, MinSubArrayLen(s, nums))
	t.Logf("nums is %v, target sum is %d, the min array len is %d", nums, s, MinSubArrayLen2(s, nums))
}

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 2}
	t.Logf("nums is %v", nums)
	tarLen := RemoveDuplicates(nums)
	t.Logf("after remove duplicates is %v", nums[:tarLen])
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	t.Logf("nums is %v", nums)
	tarLen = RemoveDuplicates(nums)
	t.Logf("after remove duplicates is %v", nums[:tarLen])
}

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	t.Logf("nums is %v", nums)
	MoveZeroes(nums)
	t.Logf("after move zero is %v", nums)
}

func BenchmarkSystemQuickSort(b *testing.B) {
	nums := []int{1, 4, 3, 2, 5, 7, 9, 2, 3, 6, 15, 24, 0, 4, 6, 20}
	for i := 0; i < b.N; i++ {
		sort.Ints(nums)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	nums := []int{1, 4, 3, 2, 5, 7, 9, 2, 3, 6, 15, 24, 0, 4, 6, 20}
	for i := 0; i < b.N; i++ {
		quickSort(nums)
	}
}

func BenchmarkQuickSortWithStack(b *testing.B) {
	nums := []int{1, 4, 3, 2, 5, 7, 9, 2, 3, 6, 15, 24, 0, 4, 6, 20}
	for i := 0; i < b.N; i++ {
		quickSortWithStack(nums)
	}
}

func BenchmarkTwoSum(b *testing.B) {
	nums := []int{2, 7, 11, 15, 25, 26, 29, 30, 35, 40, 49, 56, 57, 61, 70, 81}
	for i := 0; i < b.N; i++ {
		TwoSum(nums, 9)
	}
}

func BenchmarkTwoSum2(b *testing.B) {
	nums := []int{2, 7, 11, 15, 25, 26, 29, 30, 35, 40, 49, 56, 57, 61, 70, 81}
	for i := 0; i < b.N; i++ {
		TwoSum2(nums, 9)
	}
}

func BenchmarkMinSubArrayLen(b *testing.B) {
	nums := []int{2, 3, 1, 2, 4, 3, 8, 7, 1, 2, 5, 7, 3, 2, 1, 1, 2, 3, 1, 0, 0, 7}
	for i := 0; i < b.N; i++ {
		MinSubArrayLen(7, nums)
	}
}

func BenchmarkMinSubArrayLen2(b *testing.B) {
	nums := []int{2, 3, 1, 2, 4, 3, 8, 7, 1, 2, 5, 7, 3, 2, 1, 1, 2, 3, 1, 0, 0, 7}
	for i := 0; i < b.N; i++ {
		MinSubArrayLen2(7, nums)
	}
}