package stack

import "testing"

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