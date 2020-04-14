package array_opt

import "testing"

func TestPivotIndex(t *testing.T) {
	nums := []int{1, 7, 3, 6, 5, 6}
	t.Logf("the nums is %v, the median index is %d\n", nums, PivotIndex2(nums))
	nums = []int{1, 2, 3}
	t.Logf("the nums is %v, the median index is %d\n", nums, PivotIndex2(nums))
	nums = []int{}
	t.Logf("the nums is %v, the median index is %d\n", nums, PivotIndex2(nums))
	nums = []int{0}
	t.Logf("the nums is %v, the median index is %d\n", nums, PivotIndex2(nums))
	nums = []int{0, 0}
	t.Logf("the nums is %v, the median index is %d\n", nums, PivotIndex2(nums))
	nums = []int{0, 0, 0}
	t.Logf("the nums is %v, the median index is %d\n", nums, PivotIndex2(nums))
}

func TestDominantIndex(t *testing.T) {
	nums := []int{3, 6, 1, 0}
	t.Logf("the nums is %v, the double max num index is %d\n", nums, DominantIndex(nums))
	nums = []int{1, 2, 3, 4}
	t.Logf("the nums is %v, the double max num index is %d\n", nums, DominantIndex(nums))
}

func TestPlusOne(t *testing.T) {
	digits := []int{}
	t.Logf("the nums is %v, plus one ret is %v\n", digits, PlusOne(digits))
	digits = []int{0}
	t.Logf("the nums is %v, plus one ret is %v\n", digits, PlusOne(digits))
	digits = []int{1, 2, 3}
	t.Logf("the nums is %v, plus one ret is %v\n", digits, PlusOne(digits))
	digits = []int{9}
	t.Logf("the nums is %v, plus one ret is %v\n", digits, PlusOne(digits))
	digits = []int{9, 9, 9}
	t.Logf("the nums is %v, plus one ret is %v\n", digits, PlusOne(digits))
}

func BenchmarkPivotIndex(b *testing.B) {
	nums := []int{1, 7, 3, 6, 5, 6, 8, 79, -10, 54, 38, 8, 79, 1, 648}
	for i := 0; i < b.N; i++ {
		PivotIndex(nums)
	}
}

func BenchmarkPivotIndex2(b *testing.B) {
	nums := []int{1, 7, 3, 6, 5, 6, 8, 79, -10, 54, 38, 8, 79, 1, 648}
	for i := 0; i < b.N; i++ {
		PivotIndex2(nums)
	}
}