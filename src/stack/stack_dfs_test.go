package stack

import (
    "math/rand"
    "testing"
)

func TestMinStack(t *testing.T) {
    ms := Constructor()
    ms.Push(-2)
    ms.Push(0)
    ms.Push(-3)
    t.Logf("the min is %d\n", ms.GetMin())
    ms.Pop()
    t.Logf("the top is %d\n", ms.Top())
    t.Logf("the min is %d\n", ms.GetMin())
}

func TestIsValid(t *testing.T) {
    t.Logf("the string is (), the ret is %t\n", IsValid("()"))
    t.Logf("the string is ()[]{}, the ret is %t\n", IsValid("()[]{}"))
    t.Logf("the string is (], the ret is %t\n", IsValid("(]"))
    t.Logf("the string is ([)], the ret is %t\n", IsValid("([)]"))
    t.Logf("the string is {[]}, the ret is %t\n", IsValid("{[]}"))
    t.Logf("the string is {[], the ret is %t\n", IsValid("{[]"))
    t.Logf("the string is {, the ret is %t\n", IsValid("{"))
    t.Logf("the string is ], the ret is %t\n", IsValid("]"))
    t.Logf("the string is empty, the ret is %t\n", IsValid(""))
}

func TestDailyTemperatures(t *testing.T) {
    temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
    t.Log(DailyTemperatures(temperatures))
}

func TestDailyTempWithoutStack(t *testing.T) {
    temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
    t.Log(DailyTempWithoutStack(temperatures))
}

func TestEvalRPNDesc(t *testing.T) {
    tokens := []string{"2", "1", "+", "3", "*"}
    t.Logf("the expression is %v, the result is %d\n", tokens, EvalRPNDesc(tokens))
    tokens = []string{"4", "13", "5", "/", "+"}
    t.Logf("the expression is %v, the result is %d\n", tokens, EvalRPNDesc(tokens))
    tokens = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
    t.Logf("the expression is %v, the result is %d\n", tokens, EvalRPNDesc(tokens))
}

func TestEvalRPNAsc(t *testing.T) {
    tokens := []string{"2", "1", "+", "3", "*"}
    t.Logf("the expression is %v, the result is %d\n", tokens, EvalRPNAsc(tokens))
    tokens = []string{"4", "13", "5", "/", "+"}
    t.Logf("the expression is %v, the result is %d\n", tokens, EvalRPNAsc(tokens))
    tokens = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
    t.Logf("the expression is %v, the result is %d\n", tokens, EvalRPNAsc(tokens))
}

func TestNumIslands(t *testing.T) {
    grid := [][]byte {
        {'1', '1', '1', '1', '0'},
        {'1', '1', '0', '1', '0'},
        {'1', '1', '0', '0', '1'},
        {'1', '1', '0', '1', '1'},
        {'0', '0', '1', '0', '1'},
    }
    t.Logf("the num of island is %d\n", NumIslandsDFS(grid))
}

func TestCloneGraph(t *testing.T) {
    node1 := &Node{Val: 1}
    node2 := &Node{Val: 2}
    node3 := &Node{Val: 3}
    node4 := &Node{Val: 4}
    node1.Neighbors = []*Node{node2, node4}
    node2.Neighbors = []*Node{node1, node3}
    node3.Neighbors = []*Node{node2, node4}
    node4.Neighbors = []*Node{node1, node3}

    t.Logf("node1 is %v\n", node1)
    t.Logf("node2 is %v\n", node1.Neighbors[0])
    t.Logf("node3 is %v\n", node1.Neighbors[0].Neighbors[1])
    t.Logf("node4 is %v\n", node1.Neighbors[1])

    cloneNode1 := CloneGraph1(node1)
    t.Logf("clone method 1:")
    t.Logf("clone node1 addr is %p\n", cloneNode1)
    t.Logf("clone node2 addr is %p\n", cloneNode1.Neighbors[0])
    t.Logf("clone node3 addr is %p\n", cloneNode1.Neighbors[0].Neighbors[1])
    t.Logf("clone node4 addr is %p\n", cloneNode1.Neighbors[1])
    t.Logf("clone node1 is %v\n", cloneNode1)
    t.Logf("clone node2 is %v\n", cloneNode1.Neighbors[0])
    t.Logf("clone node3 is %v\n", cloneNode1.Neighbors[0].Neighbors[1])
    t.Logf("clone node4 is %v\n", cloneNode1.Neighbors[1])

    cloneNode1 = CloneGraph1(node1)
    t.Logf("clone method 2:")
    t.Logf("clone node1 addr is %p\n", cloneNode1)
    t.Logf("clone node2 addr is %p\n", cloneNode1.Neighbors[0])
    t.Logf("clone node3 addr is %p\n", cloneNode1.Neighbors[0].Neighbors[1])
    t.Logf("clone node4 addr is %p\n", cloneNode1.Neighbors[1])
    t.Logf("clone node1 is %v\n", cloneNode1)
    t.Logf("clone node2 is %v\n", cloneNode1.Neighbors[0])
    t.Logf("clone node3 is %v\n", cloneNode1.Neighbors[0].Neighbors[1])
    t.Logf("clone node4 is %v\n", cloneNode1.Neighbors[1])
}

func BenchmarkDailyTempWithoutStack(b *testing.B) {
    arrLen := 30000
    temperatures := make([]int, arrLen)
    for i := 0; i < arrLen; i++ {
        temperatures[i] = 30 + rand.Intn(70)
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        DailyTempWithoutStack(temperatures)
    }
}

func BenchmarkDailyTemperatures(b *testing.B) {
    arrLen := 30000
    temperatures := make([]int, arrLen)
    for i := 0; i < arrLen; i++ {
        temperatures[i] = 30 + rand.Intn(70)
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        DailyTemperatures(temperatures)
    }
}

func BenchmarkEvalRPNDesc(b *testing.B) {
    tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        EvalRPNDesc(tokens)
    }
}

func BenchmarkEvalRPNAsc(b *testing.B) {
    tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        EvalRPNAsc(tokens)
    }
}

func BenchmarkNumIslandsDFS(b *testing.B) {
    grid := [][]byte {
        {'1', '1', '1', '1', '0'},
        {'1', '1', '0', '1', '0'},
        {'1', '1', '0', '0', '1'},
        {'1', '1', '0', '1', '1'},
        {'0', '0', '1', '0', '1'},
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        NumIslandsDFS(grid)
    }
}

func BenchmarkCloneGraph1(b *testing.B) {
    node1 := &Node{Val: 1}
    node2 := &Node{Val: 2}
    node3 := &Node{Val: 3}
    node4 := &Node{Val: 4}
    node1.Neighbors = []*Node{node2, node4}
    node2.Neighbors = []*Node{node1, node3}
    node3.Neighbors = []*Node{node2, node4}
    node4.Neighbors = []*Node{node1, node3}
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        CloneGraph1(node1)
    }
}

func BenchmarkCloneGraph2(b *testing.B) {
    node1 := &Node{Val: 1}
    node2 := &Node{Val: 2}
    node3 := &Node{Val: 3}
    node4 := &Node{Val: 4}
    node1.Neighbors = []*Node{node2, node4}
    node2.Neighbors = []*Node{node1, node3}
    node3.Neighbors = []*Node{node2, node4}
    node4.Neighbors = []*Node{node1, node3}
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        CloneGraph2(node1)
    }
}