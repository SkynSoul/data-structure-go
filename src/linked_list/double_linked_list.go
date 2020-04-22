package linked_list

import "fmt"

type DoubleListNode struct {
    Val int
    Prev *DoubleListNode
    Next *DoubleListNode
}

type DoubleLinkedList struct {
    head *DoubleListNode
    tail *DoubleListNode
    length int
}

func ConstructorDoubleList() DoubleLinkedList {
    headNode, tailNode := &DoubleListNode{Val: -1, Prev: nil, Next: nil}, &DoubleListNode{Val: -1, Prev: nil, Next: nil}
    headNode.Next = tailNode
    tailNode.Prev = headNode
    return DoubleLinkedList{
        head: headNode,
        tail: tailNode,
        length: 0,
    }
}

func (dl *DoubleLinkedList) getNodeByIndex(index int) *DoubleListNode {
    if index < 0 || index >= dl.length {
        return nil
    }
    curNode := dl.head
    if index < dl.length - index {
        for i := 0; i <= index; i++ {
            curNode = curNode.Next
        }
    } else {
        curNode = dl.tail
        for i := 0; i < dl.length - index; i++ {
            curNode = curNode.Prev
        }
    }
    return curNode
}

func (dl *DoubleLinkedList) Get(index int) int {
    tarNode := dl.getNodeByIndex(index)
    if tarNode == nil {
        return -1
    }
    return tarNode.Val
}

func (dl *DoubleLinkedList) AddAtHead(val int)  {
    inNode := &DoubleListNode{Val: val, Prev: dl.head, Next: dl.head.Next}
    dl.head.Next.Prev = inNode
    dl.head.Next = inNode
    dl.length++
}

func (dl *DoubleLinkedList) AddAtTail(val int)  {
    inNode := &DoubleListNode{Val: val, Prev: dl.tail.Prev, Next: dl.tail}
    dl.tail.Prev.Next = inNode
    dl.tail.Prev = inNode
    dl.length++
}

func (dl *DoubleLinkedList) AddAtIndex(index int, val int)  {
    if index <= 0 {
        dl.AddAtHead(val)
        return
    }
    if index == dl.length {
        dl.AddAtTail(val)
        return
    }
    if index > dl.length {
        return
    }
    nextNode := dl.getNodeByIndex(index)
    inNode := &DoubleListNode{Val: val, Prev: nextNode.Prev, Next: nextNode}
    nextNode.Prev.Next = inNode
    nextNode.Prev = inNode
    dl.length++
}

func (dl *DoubleLinkedList) DeleteAtIndex(index int)  {
    delNode := dl.getNodeByIndex(index)
    if delNode == nil {
        return
    }
    delNode.Prev.Next = delNode.Next
    delNode.Next.Prev = delNode.Prev
    dl.length--
}

func (dl *DoubleLinkedList) Print() {
    fmt.Printf("-----------------\nthe list length is %d\n", dl.length)
    curNode := dl.head
    for i := 0; i < dl.length; i++ {
        if curNode.Next == nil {
            return
        }
        curNode = curNode.Next
        fmt.Printf("%d\t", curNode.Val)
    }
    fmt.Printf("\n")
}