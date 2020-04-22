package linked_list

import (
    "fmt"
    "testing"
)

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

func TestGetIntersectionNode(t *testing.T) {
    d := &ListNode{Val: 4, Next: nil}
    c := &ListNode{Val: 3, Next: d}
    b := &ListNode{Val: 2, Next: c}
    a := &ListNode{Val: 1, Next: b}
    aa := &ListNode{Val: 0, Next: c}
    t.Logf("the intersection node is %v", GetIntersectionNode(a, aa))
    t.Logf("a: %v, aa: %v", a, aa)
    t.Logf("the intersection node is %v", GetIntersectionNode2(a, aa))
    t.Logf("a: %v, aa: %v", a, aa)
}

func TestRemoveNthFromEnd(t *testing.T) {
    a := &ListNode{Val: 1, Next: nil}
    t.Logf("after delete %d\n", 1)
    head := RemoveNthFromEnd(a, 1)
    for head != nil {
        fmt.Printf("%d\t", head.Val)
        head = head.Next
    }
    fmt.Println("")
    
    c := &ListNode{Val: 3, Next: nil}
    b := &ListNode{Val: 2, Next: c}
    a = &ListNode{Val: 1, Next: b}
    t.Logf("after delete %d\n", 1)
    head = RemoveNthFromEnd(a, 1)
    for head != nil {
        fmt.Printf("%d\t", head.Val)
        head = head.Next
    }
    fmt.Println("")
    
    c = &ListNode{Val: 3, Next: nil}
    b = &ListNode{Val: 2, Next: c}
    a = &ListNode{Val: 1, Next: b}
    t.Logf("after delete %d\n", 3)
    head = RemoveNthFromEnd(a, 3)
    for head != nil {
        fmt.Printf("%d\t", head.Val)
        head = head.Next
    }
    fmt.Println("")
    
    c = &ListNode{Val: 3, Next: nil}
    b = &ListNode{Val: 2, Next: c}
    a = &ListNode{Val: 1, Next: b}
    t.Logf("after delete %d\n", 2)
    head = RemoveNthFromEnd(a, 2)
    for head != nil {
        fmt.Printf("%d\t", head.Val)
        head = head.Next
    }
    fmt.Println("")
}

func TestReverseList(t *testing.T) {
    c := &ListNode{Val: 3, Next: nil}
    b := &ListNode{Val: 2, Next: c}
    a := &ListNode{Val: 1, Next: b}
    t.Logf("after reverse")
    head := ReverseList(a)
    for head != nil {
        fmt.Printf("%d\t", head.Val)
        head = head.Next
    }
    fmt.Println("")
    
    a = &ListNode{Val: 1, Next: nil}
    t.Logf("after reverse")
    head = ReverseList(a)
    for head != nil {
        fmt.Printf("%d\t", head.Val)
        head = head.Next
    }
    fmt.Println("")
}

func TestReverseList2(t *testing.T) {
    e := &ListNode{Val: 5, Next: nil}
    d := &ListNode{Val: 4, Next: e}
    c := &ListNode{Val: 3, Next: d}
    b := &ListNode{Val: 2, Next: c}
    a := &ListNode{Val: 1, Next: b}
    t.Logf("after reverse")
    head := ReverseList2(a)
    for head != nil {
        fmt.Printf("%d\t", head.Val)
        head = head.Next
    }
    fmt.Println("")
    
    a = &ListNode{Val: 1, Next: nil}
    t.Logf("after reverse")
    head = ReverseList2(a)
    for head != nil {
        fmt.Printf("%d\t", head.Val)
        head = head.Next
    }
    fmt.Println("")
}

func TestReverseList3(t *testing.T) {
    e := &ListNode{Val: 5, Next: nil}
    d := &ListNode{Val: 4, Next: e}
    c := &ListNode{Val: 3, Next: d}
    b := &ListNode{Val: 2, Next: c}
    a := &ListNode{Val: 1, Next: b}
    t.Logf("after reverse")
    head := ReverseList3(a)
    for head != nil {
        fmt.Printf("%d\t", head.Val)
        head = head.Next
    }
    fmt.Println("")
    
    a = &ListNode{Val: 1, Next: nil}
    t.Logf("after reverse")
    head = ReverseList3(a)
    for head != nil {
        fmt.Printf("%d\t", head.Val)
        head = head.Next
    }
    fmt.Println("")
}

func TestRemoveElements(t *testing.T) {
    l := ConstructorSinglyList()
    for i := 1; i <= 6; i++ {
        l.AddAtTail(i)
        l.AddAtTail(i)
    }
    l.Print()
    head := RemoveElements(l.head.Next, 6)
    for head != nil {
        t.Logf("%d", head.Val)
        head = head.Next
    }
}

func TestOddEvenList(t *testing.T) {
    l := ConstructorSinglyList()
    for i := 1; i <= 3; i++ {
        l.AddAtTail(i)
    }
    l.Print()
    head := OddEvenList(l.head.Next)
    for head != nil {
        t.Logf("%d", head.Val)
        head = head.Next
    }
}

func TestIsPalindrome(t *testing.T) {
    l := ConstructorSinglyList()
    for i := 1; i <= 2; i++ {
       l.AddAtTail(i)
    }
    //for i := 2; i > 0; i-- {
    //   l.AddAtTail(i)
    //}
    l.Print()
    t.Logf("the list is palindrome: %t\n", IsPalindrome(l.head.Next))
}