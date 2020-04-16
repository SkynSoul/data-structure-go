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
