package stack

type MinStack struct {
    data []int
    min []int
}


func Constructor() MinStack {
    return MinStack{data: make([]int, 0, 5), min: make([]int, 0, 5)}
}


func (ms *MinStack) Push(x int)  {
    ms.data = append(ms.data, x)
    minLen := len(ms.min)
    if minLen > 0 && ms.min[minLen - 1] < x {
        x = ms.min[minLen - 1]
    }
    ms.min = append(ms.min, x)
}


func (ms *MinStack) Pop()  {
    if len(ms.data) > 0 {
        ms.data = ms.data[:len(ms.data) - 1]
        ms.min = ms.min[:len(ms.min) - 1]
    }
}


func (ms *MinStack) Top() int {
    msLen := len(ms.data)
    if msLen == 0 {
        return -1
    }
    return ms.data[msLen - 1]
}


func (ms *MinStack) GetMin() int {
    msLen := len(ms.min)
    if msLen == 0 {
        return -1
    }
    return ms.min[msLen - 1]
}
