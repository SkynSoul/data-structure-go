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