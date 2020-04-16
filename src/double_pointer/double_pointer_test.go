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