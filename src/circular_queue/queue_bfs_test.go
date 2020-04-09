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
