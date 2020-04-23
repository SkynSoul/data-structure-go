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

func TestDoubleLinkedList(t *testing.T) {
    l := ConstructorDoubleList()
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

func TestMergeTwoLists(t *testing.T) {
    l1 := ConstructorSinglyList()
    l1.AddAtTail(0)
    l1.AddAtTail(0)
    l1.AddAtTail(1)
    l1.AddAtTail(5)
    l2 := ConstructorSinglyList()
    l2.AddAtTail(1)
    l2.AddAtTail(2)
    l2.AddAtTail(4)
    mHead := MergeTwoLists(l1.head.Next, l2.head.Next)
    for mHead != nil {
        t.Logf("%d", mHead.Val)
        mHead = mHead.Next
    }
}

func TestAddTwoNumbers(t *testing.T) {
    l1 := ConstructorSinglyList()
    l1.AddAtTail(2)
    l1.AddAtTail(4)
    l1.AddAtTail(3)
    l2 := ConstructorSinglyList()
    l2.AddAtTail(5)
    l2.AddAtTail(6)
    l2.AddAtTail(4)
    mHead := AddTwoNumbers(l1.head.Next, l2.head.Next)
    for mHead != nil {
        t.Logf("%d", mHead.Val)
        mHead = mHead.Next
    }
}

func TestFlatten(t *testing.T) {
    nodeArr := make([]*Node, 12)
    for i := 0; i < 12; i++ {
        nodeArr[i] = &Node{Val: i + 1}
    }
    for i := 0; i < 12; i++ {
        if i != 0 {
            nodeArr[i].Prev = nodeArr[i - 1]
        }
        if i != 11 {
            nodeArr[i].Next = nodeArr[i + 1]
        }
    }
    nodeArr[5].Next = nil
    nodeArr[9].Next = nil
    nodeArr[2].Child = nodeArr[6]
    nodeArr[6].Prev = nil
    nodeArr[7].Child = nodeArr[10]
    nodeArr[10].Prev = nil
    
    root := Flatten(nodeArr[0])
    for root != nil {
        t.Logf("%v", root)
        root = root.Next
    }
}

func TestCopyRandomList(t *testing.T) {
    n5 := &RandomNode{Val: 1, Next: nil}
    n4 := &RandomNode{Val: 10, Next: n5}
    n3 := &RandomNode{Val: 11, Next: n4}
    n2 := &RandomNode{Val: 13, Next: n3}
    n1 := &RandomNode{Val: 7, Next: n2}
    n2.Random = n1
    n3.Random = n5
    n4.Random = n3
    n5.Random = n1
    nn1 := CopyRandomList(n1)
    for nn1 != nil {
        t.Logf("addr: %p, val: %d, struct: %v", nn1, nn1.Val, nn1)
        nn1 = nn1.Next
    }
}

func TestCopyRandomListByStack(t *testing.T) {
    n5 := &RandomNode{Val: 1, Next: nil}
    n4 := &RandomNode{Val: 10, Next: n5}
    n3 := &RandomNode{Val: 11, Next: n4}
    n2 := &RandomNode{Val: 13, Next: n3}
    n1 := &RandomNode{Val: 7, Next: n2}
    n2.Random = n1
    n3.Random = n5
    n4.Random = n3
    n5.Random = n1
    nn1 := CopyRandomListByStack(n1)
    for nn1 != nil {
        t.Logf("addr: %p, val: %d, struct: %v", nn1, nn1.Val, nn1)
        nn1 = nn1.Next
    }
}

func BenchmarkCopyRandomList(b *testing.B) {
    n5 := &RandomNode{Val: 1, Next: nil}
    n4 := &RandomNode{Val: 10, Next: n5}
    n3 := &RandomNode{Val: 11, Next: n4}
    n2 := &RandomNode{Val: 13, Next: n3}
    n1 := &RandomNode{Val: 7, Next: n2}
    n2.Random = n1
    n3.Random = n5
    n4.Random = n3
    n5.Random = n1
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        CopyRandomList(n1)
    }
}

func BenchmarkCopyRandomListByStack(b *testing.B) {
    n5 := &RandomNode{Val: 1, Next: nil}
    n4 := &RandomNode{Val: 10, Next: n5}
    n3 := &RandomNode{Val: 11, Next: n4}
    n2 := &RandomNode{Val: 13, Next: n3}
    n1 := &RandomNode{Val: 7, Next: n2}
    n2.Random = n1
    n3.Random = n5
    n4.Random = n3
    n5.Random = n1
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        CopyRandomListByStack(n1)
    }
}