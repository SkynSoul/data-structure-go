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