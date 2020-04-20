package linked_list

import "testing"

func TestSinglyLinkedList(t *testing.T) {
    l := ConstructorSinglyList()
    l.Print()
    t.Logf("l get(1): %d", l.Get(1))
    l.AddAtHead(1)
    l.Print()
    l.AddAtTail(3)
    l.Print()
    l.AddAtIndex(1, 2)
    l.Print()
    t.Logf("l get(2): %d", l.Get(2))
    t.Logf("l get(1): %d", l.Get(1))
    l.DeleteAtIndex(1)
    l.Print()
    t.Logf("l get(1): %d", l.Get(1))
    l.DeleteAtIndex(0)
    l.Print()
    l.AddAtIndex(0, 2)
    l.Print()
    l.AddAtIndex(0, 1)
    l.Print()
    l.AddAtIndex(-1, 0)
    l.Print()
    l.AddAtIndex(100, 4)
    l.Print()
    l.AddAtIndex(5, 5)
    l.Print()
}

func TestHasCycle(t *testing.T) {
    d := &ListNode{Val: 4, Next: nil}
    c := &ListNode{Val: 3, Next: d}
    b := &ListNode{Val: 2, Next: c}
    a := &ListNode{Val: 1, Next: b}
    t.Logf("has cycle: %t", HasCycle(a))
    t.Logf("detect cycle: %v", DetectCycle(a))
    d.Next = c
    t.Logf("has cycle: %t", HasCycle(a))
    t.Logf("detect cycle: %v", DetectCycle(a))
}