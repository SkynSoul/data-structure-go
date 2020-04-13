package queue_stack

type MyStack struct {
	innerQueue *InnerQueue
}

/** Initialize your data structure here. */
func ConstructorStack() MyStack {
	return MyStack{innerQueue: NewInnerQueue()}
}

/** Push element x onto stack. */
func (ms *MyStack) Push(x int)  {
	ms.innerQueue.Push(x)
}

/** Removes the element on top of the stack and returns that element. */
func (ms *MyStack) Pop() int {
	if ms.innerQueue.IsEmpty() {
		return -1
	}
	tempQueue := NewInnerQueue()
	for ms.innerQueue.Size() > 1 {
		tempQueue.Push(ms.innerQueue.Pop())
	}
	ret := ms.innerQueue.Pop().(int)
	for !tempQueue.IsEmpty() {
		ms.innerQueue.Push(tempQueue.Pop())
	}
	return ret
}

/** Get the top element. */
func (ms *MyStack) Top() int {
	if ms.innerQueue.IsEmpty() {
		return -1
	}
	tempQueue := NewInnerQueue()
	for ms.innerQueue.Size() > 1 {
		tempQueue.Push(ms.innerQueue.Pop())
	}
	tempQueue.Push(ms.innerQueue.Peek())
	ret := ms.innerQueue.Pop().(int)
	for !tempQueue.IsEmpty() {
		ms.innerQueue.Push(tempQueue.Pop())
	}
	return ret
}

/** Returns whether the stack is empty. */
func (ms *MyStack) Empty() bool {
	return ms.innerQueue.IsEmpty()
}