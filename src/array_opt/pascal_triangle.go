package array_opt

func GeneratePascalTriangle(numRows int) [][]int  {
    ret := make([][]int, 0, numRows)
    if numRows == 0 {
        return ret
    }
    ret = append(ret, []int{1})
    for x := 1; x < numRows; x++ {
        row := make([]int, x + 1)
        for y := 0; y <= x; y++ {
            row[y] = getVal(ret[x - 1], y - 1) + getVal(ret[x - 1], y)
        }
        ret = append(ret, row)
    }
    return ret
}

func getVal(set []int, idx int) int {
    if idx < 0 || idx >= len(set) {
        return 0
    }
    return set[idx]
}

func GetRow(rowIndex int) []int {
    if rowIndex < 0 {
        return []int{}
    }
    return generateRow(1, rowIndex + 1, []int{})
}

func generateRow(rowNum int, target int, lastRow []int) []int {
    curRow := make([]int, rowNum)
    curRow[0] = 1
    for i := 1; i < rowNum; i++ {
        curRow[i] = getVal(lastRow, i - 1) + getVal(lastRow, i)
    }
    if rowNum == target {
        return curRow
    }
    rowNum++
    return generateRow(rowNum, target, curRow)
}

// 计算公式 n = 当前行索引 k = 当前列索引  从0开始
// n的第k个元素 = n! / (k! * (n - k)!)
func GetRow2(rowIndex int) []int {
    ret := make([]int, rowIndex + 1)
    for i := 0; i <= rowIndex; i++ {
        ret[i] = calFactorial(rowIndex) / (calFactorial(i) * calFactorial(rowIndex - i))
    }
    return ret
}

func calFactorial(n int) int {
    ret := 1
    for i := 1; i <= n; i++ {
        ret *= i
    }
    return ret
}

func GetRow3(rowIndex int) []int {
    preRow := make([]int, rowIndex + 1)
    curRow := make([]int, rowIndex + 1)
    preRow[0] = 1
    for i := 0; i <= rowIndex; i++ {
        for j := 0; j <= i; j++ {
            curRow[j] = getVal(preRow, j - 1) + getVal(preRow, j)
        }
        preRow, curRow = curRow, preRow
    }
    return preRow
}