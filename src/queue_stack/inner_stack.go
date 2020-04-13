package queue_stack

import "container/list"

type InnerStack struct {
	data *list.List
}

func NewInnerStack() *InnerStack {
	return &InnerStack{data: list.New()}
}

func (is *InnerStack) Push(val interface{}) {
	is.data.PushBack(val)
}

func (is *InnerStack) Pop() interface{} {
	if is.IsEmpty() {
		return nil
	}
	return is.data.Remove(is.data.Back())
}

func (is *InnerStack) Top() interface{} {
	if is.IsEmpty() {
		return nil
	}
	return is.data.Back().Value
}

func (is *InnerStack) IsEmpty() bool {
	return is.data.Len() == 0
}

func (is *InnerStack) Size() int {
	return is.data.Len()
}