package queue_stack

import "container/list"

type InnerQueue struct {
	data *list.List
}

func NewInnerQueue() *InnerQueue  {
	return &InnerQueue{data: list.New()}
}

func (iq *InnerQueue) Push(val interface{}) {
	iq.data.PushBack(val)
}

func (iq *InnerQueue) Pop() interface{} {
	if iq.IsEmpty() {
		return nil
	}
	return iq.data.Remove(iq.data.Front())
}

func (iq *InnerQueue) Peek() interface{} {
	if iq.IsEmpty() {
		return nil
	}
	return iq.data.Front().Value
}

func (iq *InnerQueue) IsEmpty() bool {
	return iq.data.Len() == 0
}

func (iq *InnerQueue) Size() int {
	return iq.data.Len()
}