package circular_queue

import "math"

// 广度优先解法
func NumSquares(n int) int {
    if n < 1 {
        return 0
    }
    // 计算可用完全平方数
    numArr := make([]int, 0)
    for i := 1; i <= n; i++ {
        if isPerfectSquare(i) {
            numArr = append(numArr, i)
        }
    }
    numLen := len(numArr)
    
    nQueue := Constructor(int(math.Pow(float64(numLen), 2)))
    step := 0
    nQueue.EnQueue(n)
    for !nQueue.IsEmpty() {
        curSize := nQueue.Size()
        for i := 0; i < curSize; i++ {
            curValue := nQueue.Front()
            nQueue.DeQueue()
            if curValue == 0 {
                return step
            }
            for j := numLen - 1; j >= 0; j-- {
                if numArr[j] > curValue {
                    continue
                }
                nQueue.EnQueue(curValue - numArr[j])
            }
        }
        step++
    }
    return -1
}

func isPerfectSquare(value int) bool {
    if value <= 0 {
        return false
    }
    var ret float64 = math.Sqrt(float64(value))
    return ret == math.Floor(ret)
}
