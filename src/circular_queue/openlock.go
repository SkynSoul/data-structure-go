package circular_queue

import "fmt"

func OpenLock(deadends []string, target string) int {
    alreadyMap := make(map[string]bool)
    for _, str := range deadends {
        alreadyMap[str] = true
    }
    if target == "0000" {
        return -1
    }
    if _, ok := alreadyMap[target]; ok {
        return -1
    }
    
    sQueue := Constructor(8192)
    step := 0
    idxEnQueue4(&sQueue, alreadyMap, 0, 0, 0, 0)
    for !sQueue.IsEmpty() {
        qSize := sQueue.Size()
        //fmt.Println(qSize)
        for i := 0; i + 3 < qSize; i += 4 {
            a, b, c, d := idxDeQueue4(&sQueue)
            curStr, _ := getNumStr(a, b, c, d)
            if curStr == target {
                return step
            }
            idxEnQueue4(&sQueue, alreadyMap, getFixIdx(a, 1), b, c, d)
            idxEnQueue4(&sQueue, alreadyMap, a, getFixIdx(b, 1), c, d)
            idxEnQueue4(&sQueue, alreadyMap, a, b, getFixIdx(c, 1), d)
            idxEnQueue4(&sQueue, alreadyMap, a, b, c, getFixIdx(d, 1))
            idxEnQueue4(&sQueue, alreadyMap, getFixIdx(a, -1), b, c, d)
            idxEnQueue4(&sQueue, alreadyMap, a, getFixIdx(b, -1), c, d)
            idxEnQueue4(&sQueue, alreadyMap, a, b, getFixIdx(c, -1), d)
            idxEnQueue4(&sQueue, alreadyMap, a, b, c, getFixIdx(d, -1))
        }
        step++
    }
    return -1
}

func getNumStr(a, b, c, d int) (string, bool) {
    if !isValid(a) || !isValid(b) || !isValid(c) || !isValid(d) {
        return "", false
    }
    lockArr := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
    return fmt.Sprintf("%c%c%c%c", lockArr[a], lockArr[b], lockArr[c], lockArr[d]), true
}

func isValid(idx int) bool {
    return idx >= 0 && idx <= 9
}

func getFixIdx(srcIdx, value int) int {
    return (srcIdx + value + 10) % 10
}

func idxEnQueue4(sQueue *MyCircularQueue, alreadyMap map[string]bool, a, b, c, d int) bool {
    str, ok := getNumStr(a, b, c, d)
    if !ok {
        return false
    }
    if _, ok = alreadyMap[str]; ok {
        return false
    }
    if sQueue.EnQueue(a) && sQueue.EnQueue(b) && sQueue.EnQueue(c) && sQueue.EnQueue(d) {
        alreadyMap[str] = true
        return true
    }
    return false
}

func idxDeQueue4(sQueue *MyCircularQueue) (a, b, c, d int) {
    a = sQueue.Front()
    sQueue.DeQueue()
    b = sQueue.Front()
    sQueue.DeQueue()
    c = sQueue.Front()
    sQueue.DeQueue()
    d = sQueue.Front()
    sQueue.DeQueue()
    return
}