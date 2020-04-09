package circular_queue

type MyCircularQueue struct {
    head int
    rear int
    size int
    data []int
}


/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{0, 0, k + 1, make([]int, k + 1)}
}


/** Insert an element into the circular queue. Return true if the operation is successful. */
func (cq *MyCircularQueue) EnQueue(value int) bool {
    if cq.IsFull() {
        return false
    }
    cq.data[cq.rear] = value
    cq.rear = (cq.rear + 1) % cq.size
    return true
}


/** Delete an element from the circular queue. Return true if the operation is successful. */
func (cq *MyCircularQueue) DeQueue() bool {
    if cq.IsEmpty() {
        return false
    }
    cq.head = (cq.head + 1) % cq.size
    return true
}


/** Get the front item from the queue. */
func (cq *MyCircularQueue) Front() int {
    if cq.IsEmpty() {
        return -1
    }
    return cq.data[cq.head]
}


/** Get the last item from the queue. */
func (cq *MyCircularQueue) Rear() int {
    if cq.IsEmpty() {
        return -1
    }
    return cq.data[(cq.rear + cq.size - 1) % cq.size]
}


/** Checks whether the circular queue is empty or not. */
func (cq *MyCircularQueue) IsEmpty() bool {
    return cq.head == cq.rear
}


/** Checks whether the circular queue is full or not. */
func (cq *MyCircularQueue) IsFull() bool {
    return (cq.rear+ 1) % cq.size == cq.head
}

func (cq *MyCircularQueue) Size() int {
    return (cq.rear - cq.head + cq.size) % cq.size
}
