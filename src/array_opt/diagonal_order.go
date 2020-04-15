package array_opt

func FindDiagonalOrder(matrix [][]int) []int {
    xMax := len(matrix)
    yMax := 0
    for i := 0; i < xMax; i++ {
        if len(matrix[i]) > yMax {
            yMax = len(matrix[i])
        }
    }
    ret := make([]int, 0)
    step := 0
    x, y := 0, 0
    for x < xMax || y < yMax {
        if x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[x]) {
            ret = append(ret, matrix[x][y])
        }
        nextX, nextY := x - 1, y + 1
        if step % 2 != 0 {
            nextX, nextY = x + 1, y - 1
        }
        if nextX < 0 || nextY < 0 {
            if nextX < 0 {
                nextX = 0
            } else {
                nextY = 0
            }
            step++
        }
        x, y = nextX, nextY
    }
    return ret
}

func FindDiagonalOrder2(matrix [][]int) []int {
    xMax := len(matrix)
    yMax := 0
    for i := 0; i < xMax; i++ {
        if len(matrix[i]) > yMax {
            yMax = len(matrix[i])
        }
    }
    ret := make([]int, 0, xMax * yMax)
    step := 0
    x, y := 0, 0
    for x < xMax || y < yMax {
        if x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[x]) {
            ret = append(ret, matrix[x][y])
        }
        nextX, nextY := x - 1, y + 1
        if step % 2 != 0 {
            nextX, nextY = x + 1, y - 1
        }
        if nextX < 0 || nextY < 0 {
            if nextX < 0 {
                nextX = 0
            } else {
                nextY = 0
            }
            step++
        }
        x, y = nextX, nextY
    }
    return ret
}

func FindDiagonalOrder3(matrix [][]int) []int {
    xMax := len(matrix)
    yMax := 0
    for i := 0; i < xMax; i++ {
        if len(matrix[i]) > yMax {
            yMax = len(matrix[i])
        }
    }
    ret := make([]int, 0, xMax * yMax)
    step := 0
    x, y := 0, 0
    for x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[x]) {
        ret = append(ret, matrix[x][y])
        nextX, nextY := x - 1, y + 1
        if step % 2 != 0 {
            nextX, nextY = x + 1, y - 1
        }
        if nextX < 0 || nextY < 0 {
            if nextX < 0 {
                nextX = 0
            } else {
                nextY = 0
            }
            step++
        }
        x, y = nextX, nextY
    }
    return ret
}
