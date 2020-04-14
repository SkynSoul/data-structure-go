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

func TestDecodeString(t *testing.T) {
	src := "3[a]2[bc]"
	t.Logf("src is %s, ret is %s\n", src, DecodeString(src))
	src = "3[a2[c]]"
	t.Logf("src is %s, ret is %s\n", src, DecodeString(src))
	src = "2[abc]3[cd]ef"
	t.Logf("src is %s, ret is %s\n", src, DecodeString(src))
	src = "15[a]10[bc]"
	t.Logf("src is %s, ret is %s\n", src, DecodeString(src))
	src = "3[z]2[2[y]pq4[2[jk]e1[f]]]ef"
	t.Logf("src is %s, ret is %s\n", src, DecodeString(src))
}

func TestFloodFill(t *testing.T) {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	t.Logf("after image render: %v\n", FloodFill(image, 1, 1, 2))
}

func TestUpdateMatrix(t *testing.T) {
	matrix := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	t.Logf("update matrix: %v\n", UpdateMatrix2(matrix))
	
	matrix = [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}
	t.Logf("update matrix: %v\n", UpdateMatrix2(matrix))
}

func BenchmarkUpdateMatrix(b *testing.B) {
	matrix := [][]int{
		{0, 0, 0, 1, 1},
		{0, 1, 0, 1, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 1, 1, 0},
		{1, 1, 0, 1, 0},
	}
	for i := 0; i < b.N; i++ {
		UpdateMatrix(matrix)
	}
}

func BenchmarkUpdateMatrix2(b *testing.B) {
	matrix := [][]int{
		{0, 0, 0, 1, 1},
		{0, 1, 0, 1, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 1, 1, 0},
		{1, 1, 0, 1, 0},
	}
	for i := 0; i < b.N; i++ {
		UpdateMatrix2(matrix)
	}
}