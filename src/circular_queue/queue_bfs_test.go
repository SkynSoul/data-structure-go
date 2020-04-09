package circular_queue

import "testing"


func TestQueueBase(t *testing.T) {
    cq := Constructor(3)
    t.Logf("the queue is empty: %t\n", cq.IsEmpty())
    t.Logf("queue size is %d\n", cq.Size())
    t.Logf("enqueue 1, ret: %t", cq.EnQueue(1))
    t.Logf("queue size is %d\n", cq.Size())
    t.Logf("enqueue 2, ret: %t", cq.EnQueue(2))
    t.Logf("enqueue 3, ret: %t", cq.EnQueue(3))
    t.Logf("queue size is %d\n", cq.Size())
    t.Logf("enqueue 4, ret: %t", cq.EnQueue(4))
    t.Logf("queue size is %d\n", cq.Size())
    t.Logf("dequeue, ret: %t", cq.DeQueue())
    t.Logf("queue size is %d\n", cq.Size())
    t.Logf("the front is %d\n", cq.Front())
    t.Logf("the rear is %d\n", cq.Rear())
    t.Logf("the queue is full: %t\n", cq.IsFull())
}

func TestNumIslands(t *testing.T) {
    grid := [][]byte {
        {'1', '1', '1', '1', '0'},
        {'1', '1', '0', '1', '0'},
        {'1', '1', '0', '0', '1'},
        {'1', '1', '0', '1', '1'},
        {'0', '0', '1', '0', '1'},
    }
    t.Logf("the num of island is %d\n", NumIslands(grid))
}

func TestOpenLock(t *testing.T) {
    deadends := []string{"0201", "0101", "0102", "1212", "2002"}
    target := "0202"
    t.Logf("the step is %d\n", OpenLock(deadends, target))
    deadends = []string{"8888"}
    target = "0009"
    t.Logf("the step is %d\n", OpenLock(deadends, target))
    deadends = []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}
    target = "8888"
    t.Logf("the step is %d\n", OpenLock(deadends, target))
    deadends = []string{"0000"}
    target = "8888"
    t.Logf("the step is %d\n", OpenLock(deadends, target))
}

func TestNumSquares(t *testing.T) {
    t.Logf("num: -1, the min step is %d\n", NumSquares(-1))
    t.Logf("num: 0, the min step is %d\n", NumSquares(0))
    t.Logf("num: 1, the min step is %d\n", NumSquares(1))
    t.Logf("num: 3, the min step is %d\n", NumSquares(3))
    t.Logf("num: 12, the min step is %d\n", NumSquares(12))
    t.Logf("num: 13, the min step is %d\n", NumSquares(13))
    t.Logf("num: 7168, the min step is %d\n", NumSquares(7168))
}