package queue_stack

type MyQueue struct {
	innerStack *InnerStack
}

/** Initialize your data structure here. */
func ConstructorQueue() MyQueue {
	return MyQueue{innerStack: NewInnerStack()}
}

/** Push element x to the back of queue. */
func (mq *MyQueue) Push(x int)  {
	mq.innerStack.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (mq *MyQueue) Pop() int {
	if mq.Empty() {
		return -1
	}
	tempStack := NewInnerStack()
	for !mq.innerStack.IsEmpty() {
		tempStack.Push(mq.innerStack.Pop())
	}
	ret := tempStack.Pop().(int)
	for !tempStack.IsEmpty() {
		mq.innerStack.Push(tempStack.Pop())
	}
	return ret
}

/** Get the front element. */
func (mq *MyQueue) Peek() int {
	if mq.Empty() {
		return -1
	}
	tempStack := NewInnerStack()
	for !mq.innerStack.IsEmpty() {
		tempStack.Push(mq.innerStack.Pop())
	}
	ret := tempStack.Top().(int)
	for !tempStack.IsEmpty() {
		mq.innerStack.Push(tempStack.Pop())
	}
	return ret
}

/** Returns whether the queue is empty. */
func (mq *MyQueue) Empty() bool {
	return mq.innerStack.IsEmpty()
}