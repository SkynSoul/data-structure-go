package circular_queue

func NumIslands(grid [][]byte) int {
    exploreQueue := Constructor(25)
    num := 0
    for x, xLen := 0, len(grid); x < xLen; x++ {
        for y, yLen := 0, len(grid[x]); y < yLen; y++ {
            if grid[x][y] != '0' {
                num++
                idxEnQueue(&exploreQueue, grid, x, y)
                for !exploreQueue.IsEmpty() {
                    curX, curY := idxDeQueue(&exploreQueue)
                    if isValidNode(curX - 1, curY, xLen, yLen) {
                        idxEnQueue(&exploreQueue, grid, curX - 1, curY)
                    }
                    if isValidNode(curX + 1, curY, xLen, yLen) {
                        idxEnQueue(&exploreQueue, grid, curX + 1, curY)
                    }
                    if isValidNode(curX, curY - 1, xLen, yLen) {
                        idxEnQueue(&exploreQueue, grid, curX, curY - 1)
                    }
                    if isValidNode(curX, curY + 1, xLen, yLen) {
                        idxEnQueue(&exploreQueue, grid, curX, curY + 1)
                    }
                }
            }
        }
    }
    return num
}

func isValidNode(x int, y int, xMax int, yMax int) bool {
    if x >= 0 && x < xMax && y >= 0 && y < yMax {
        return true
    }
    return false
}

func idxEnQueue(exploreQueue *MyCircularQueue, grid [][]byte, x int, y int) bool {
    if grid[x][y] != '0' {
        exploreQueue.EnQueue(x)
        exploreQueue.EnQueue(y)
        grid[x][y] = '0'
        return true
    }
    return false
}

func idxDeQueue(exploreQueue *MyCircularQueue) (x int, y int) {
    x = exploreQueue.Front()
    exploreQueue.DeQueue()
    y = exploreQueue.Front()
    exploreQueue.DeQueue()
    return
}