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