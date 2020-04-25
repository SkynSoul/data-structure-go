package array_opt

import (
	"testing"
	"unsafe"
)

func TestArr(t *testing.T) {
	aa := make([][]*int, 10)
	t.Logf("%v, len: %d", aa, len(aa))
	
	bb := make([]*int, 10)
	t.Logf("%v, len: %d", bb, len(bb))

	// 数组内存分配测试
	a := [5]int{1, 2, 3}
	t.Logf("arr: %v, addr: %p, a[:] %p, a[1:] %p, a[:1] %p\n",  a, &a, a[:], a[1:], a[:1])
	arrSlice := a[:]
	t.Logf("arrSlice: %v, addr: %p\n",  arrSlice, arrSlice)
	arrSlice = a[1:]
	t.Logf("arrSlice: %v, addr: %p\n",  arrSlice, arrSlice)
	arrSlice = append(a[:], 4)
	t.Logf("arrSlice: %v, addr: %p\n",  arrSlice, arrSlice)
	arrSlice = append(arrSlice, a[:]...)
	t.Logf("arrSlice: %v, addr: %p\n",  arrSlice, arrSlice)
	
	arr := []int{1, 2, 3}
	t.Logf("%p\n",  arr)
	arr = append(arr, 4)
	t.Logf("%p\n",  arr)
	arr = append([]int{0}, arr...)
	t.Logf("%p\n",  arr)
	
	arr = make([]int, 3, 5)
	arr[0], arr[1], arr[2] = 1, 2, 3
	t.Logf("%p\n",  arr)
	arr = append(arr, 4)
	t.Logf("%p\n",  arr)
	arr = append(arr, 5)
	t.Logf("%p\n",  arr)
	arr = append(arr, 6)
	t.Logf("%p\n",  arr)
	arr = append([]int{0}, arr...)
	t.Logf("%p\n",  arr)
	
	arr2 := [3]int{1, 2, 3}
	arr2Slice := arr2[:0]
	t.Logf("arr2: %v, arrSlice: %v", arr2, arr2Slice)
	arr2Slice = append(arr2Slice, 4)
	t.Logf("arr2: %v, arrSlice: %v", arr2, arr2Slice)

	// 切片以struct实现
	//type slice struct {
	//	array unsafe.Pointer
	//	len   int
	//	cap   int
	//}
	arrs := make([][]int, 0, 5)
	arrs = append(arrs, []int{1, 2, 3})
	arrs = append(arrs, []int{1, 2})
	arrs = append(arrs, []int{1, 2, 3, 4, 5})
	t.Logf("arrs[0]: %p, arrs[1]: %p, arrs[2]: %p\n", arrs[0], arrs[1], arrs[2])
	t.Logf("arrs size is %d\n", unsafe.Sizeof(arrs))

	arrs = make([][]int, 0, 10)
	srcArr := []int{1, 2, 3, 4, 5}
	arrs = append(arrs, srcArr[:])
	arrs = append(arrs, srcArr[1:])
	arrs = append(arrs, srcArr[2:])
	arrs = append(arrs, srcArr[2:3])
	arrs = append(arrs, srcArr[:3])
	t.Logf("arrs[0]: %p, arrs[1]: %p, arrs[2]: %p, arrs[3]: %p, arrs[4]: %p\n", arrs[0], arrs[1], arrs[2], arrs[3], arrs[4])
	t.Logf("arrs size is %d\n", unsafe.Sizeof(arrs))

	t.Logf("[]int size is %d\n", unsafe.Sizeof([]int{}))		// slice结构体大小
	t.Logf("[]string size is %d\n", unsafe.Sizeof([]int{}))

	srcArr = []int{1, 2, 3, 4, 5}
	t.Logf("srcArr[:2] is %v, srcArr[2:] is %v, srcArr[1:3] is %v", srcArr[:2], srcArr[2:], srcArr[1:3])

	aArr, bArr := make([]int, 5), make([]int, 3)
	t.Logf("aArr addr: %p, bArr addr: %p", aArr, bArr)
	aArr, bArr = bArr, aArr
	t.Logf("after swap, aArr addr: %p, bArr addr: %p", aArr, bArr)

	arrMap := make(map[int][]int)
	t.Logf("append nil: %v", append(arrMap[0], 1))
}

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

func TestGeneratePascalTriangle(t *testing.T) {
	numRows := 0
	t.Logf("the row's num is %d, ret is %v\n", numRows, GeneratePascalTriangle(numRows))
	numRows = 1
	t.Logf("the row's num is %d, ret is %v\n", numRows, GeneratePascalTriangle(numRows))
	numRows = 2
	t.Logf("the row's num is %d, ret is %v\n", numRows, GeneratePascalTriangle(numRows))
	numRows = 3
	t.Logf("the row's num is %d, ret is %v\n", numRows, GeneratePascalTriangle(numRows))
	numRows = 4
	t.Logf("the row's num is %d, ret is %v\n", numRows, GeneratePascalTriangle(numRows))
	numRows = 5
	t.Logf("the row's num is %d, ret is %v\n", numRows, GeneratePascalTriangle(numRows))
	numRows = 10
	retArr := GeneratePascalTriangle(numRows)
	for i := 0; i < numRows; i++ {
		t.Logf("%d\t%v", i, retArr[i])
	}
	for i := 0; i < numRows; i++ {
		t.Logf("%d\t%v", i, GetRow(i))
	}
	for i := 0; i < numRows; i++ {
		t.Logf("%d\t%v", i, GetRow2(i))
	}
	for i := 0; i < numRows; i++ {
		t.Logf("%d\t%v", i, GetRow3(i))
	}
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

func BenchmarkGetRow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetRow(30)
	}
}

func BenchmarkGetRow2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetRow2(30)
	}
}

func BenchmarkGetRow3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetRow3(30)
	}
}