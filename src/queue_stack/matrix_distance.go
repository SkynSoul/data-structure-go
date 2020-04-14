package queue_stack

// 全部搜索
func UpdateMatrix(matrix [][]int) [][]int {
    xMax := len(matrix)
    if xMax == 0 {
        return [][]int{}
    }
    yMax := len(matrix[0])
    if yMax == 0 {
        return [][]int{}
    }
    ret := make([][]int, xMax)
    for x := 0; x < len(matrix); x++ {
       row := make([]int, yMax)
       for y := 0; y < len(matrix[x]); y++ {
           row[y] = bfsUpdate(matrix, x, y)
       }
       ret[x] = row
    }
    return ret
}

func bfsUpdate(matrix [][]int, x int, y int) int {
    qHelper := make([]int, 0, 10)
    visited := make(map[int]bool)
    key := getKey(x, y)
    qHelper = append(qHelper, key)
    visited[key] = true
    step := 0
    for len(qHelper) > 0 {
        curSize := len(qHelper)
        for i := 0; i < curSize; i++ {
            key = qHelper[i]
            curX, curY := keyToPoint(key)
            if matrix[curX][curY] == 0 {
                return step
            }
            if key = isValidPoint(matrix, curX - 1, curY, visited); key >= 0 {
                qHelper = append(qHelper, key)
                visited[key] = true
            }
            if key = isValidPoint(matrix, curX + 1, curY, visited); key >= 0 {
                qHelper = append(qHelper, key)
                visited[key] = true
            }
            if key = isValidPoint(matrix, curX, curY - 1, visited); key >= 0 {
                qHelper = append(qHelper, key)
                visited[key] = true
            }
            if key = isValidPoint(matrix, curX, curY + 1, visited); key >= 0 {
                qHelper = append(qHelper, key)
                visited[key] = true
            }
        }
        qHelper = qHelper[curSize:]
        step++
    }
    return -1
}

// 从零开始
func UpdateMatrix2(matrix [][]int) [][]int {
    qHelper := make([]int, 0, 10)
    visited := make(map[int]bool)
    for x := 0; x < len(matrix); x++ {
        for y := 0; y < len(matrix[x]); y++ {
            if matrix[x][y] == 0 {
                key := getKey(x, y)
                qHelper = append(qHelper, key)
                visited[key] = true
            }
        }
    }
    step := 0
    for len(qHelper) > 0 {
        curSize := len(qHelper)
        for i := 0; i < curSize; i++ {
            curX, curY := keyToPoint(qHelper[i])
            matrix[curX][curY] = step
            if key := isValidPoint(matrix, curX - 1, curY, visited); key >= 0 {
                qHelper = append(qHelper, key)
                visited[key] = true
            }
            if key := isValidPoint(matrix, curX + 1, curY, visited); key >= 0 {
                qHelper = append(qHelper, key)
                visited[key] = true
            }
            if key := isValidPoint(matrix, curX, curY - 1, visited); key >= 0 {
                qHelper = append(qHelper, key)
                visited[key] = true
            }
            if key := isValidPoint(matrix, curX, curY + 1, visited); key >= 0 {
                qHelper = append(qHelper, key)
                visited[key] = true
            }
        }
        qHelper = qHelper[curSize:]
        step++
    }
    return matrix
}

func getKey(x int, y int) int {
    return x * 10000 + y
}

func keyToPoint(key int) (int, int) {
    x := key / 10000
    y := key % 10000
    return x, y
}

func isValidPoint(matrix [][]int, x int, y int, visited map[int]bool) int {
    if x < 0 || x >= len(matrix) || y < 0 || y >= len(matrix[x]) {
        return -1
    }
    key := getKey(x, y)
    _, ok := visited[key]
    if ok {
        return -1
    }
    return key
}