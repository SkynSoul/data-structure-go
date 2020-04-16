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

func TestFindDiagonalOrder(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	t.Logf("diagonal order1 is: %v\n", FindDiagonalOrder(matrix))
	t.Logf("diagonal order2 is: %v\n", FindDiagonalOrder2(matrix))
	t.Logf("diagonal order3 is: %v\n", FindDiagonalOrder3(matrix))
	matrix = [][]int{
		{1, 2, 3},
	}
	t.Logf("diagonal order1 is: %v\n", FindDiagonalOrder(matrix))
	t.Logf("diagonal order2 is: %v\n", FindDiagonalOrder2(matrix))
	t.Logf("diagonal order3 is: %v\n", FindDiagonalOrder3(matrix))
	// 以下两种情况  方法三有问题
	matrix = [][]int{
		{1, 2},
		{4, 5, 6},
		{7, 8, 9},
	}
	t.Logf("diagonal order1 is: %v\n", FindDiagonalOrder(matrix))
	t.Logf("diagonal order2 is: %v\n", FindDiagonalOrder2(matrix))
	t.Logf("diagonal order3 is: %v\n", FindDiagonalOrder3(matrix))
	matrix = [][]int{
		{1, 2, 3},
		{4, 5},
		{7, 8},
	}
	t.Logf("diagonal order1 is: %v\n", FindDiagonalOrder(matrix))
	t.Logf("diagonal order2 is: %v\n", FindDiagonalOrder2(matrix))
	t.Logf("diagonal order3 is: %v\n", FindDiagonalOrder3(matrix))
}

func TestSpiralOrder(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	t.Logf("spiral roder is: %v\n", SpiralOrder(matrix))
	matrix = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	t.Logf("spiral roder is: %v\n", SpiralOrder(matrix))
	matrix = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}
	t.Logf("spiral roder is: %v\n", SpiralOrder(matrix))
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

// 测试一下append的内存分配
func BenchmarkFindDiagonalOrder(b *testing.B) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for i := 0; i < b.N; i++ {
		FindDiagonalOrder(matrix)
	}
}

func BenchmarkFindDiagonalOrder2(b *testing.B) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for i := 0; i < b.N; i++ {
		FindDiagonalOrder2(matrix)
	}
}

func BenchmarkFindDiagonalOrder3(b *testing.B) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for i := 0; i < b.N; i++ {
		FindDiagonalOrder3(matrix)
	}
}