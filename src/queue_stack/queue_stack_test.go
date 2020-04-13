package queue_stack

import "testing"

func TestStackToQueue(t *testing.T) {
	myQueue := ConstructorQueue()
	myQueue.Push(1)
	myQueue.Push(2)
	t.Logf("peek returns %d\n", myQueue.Peek())
	t.Logf("pop returns %d\n", myQueue.Pop())
	t.Logf("is empty returns %t\n", myQueue.Empty())
}

func TestQueueToStack(t *testing.T) {
	myStack := ConstructorStack()
	myStack.Push(1)
	myStack.Push(2)
	t.Logf("pop returns %d\n", myStack.Pop())
	t.Logf("top returns %d\n", myStack.Top())
	t.Logf("is empty returns %t\n", myStack.Empty())
}