package linked_list

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

type MyLinkedList struct {
    head   *ListNode
    length int
}

/** Initialize your data structure here. */
func ConstructorSinglyList() MyLinkedList {
    return MyLinkedList{head: &ListNode{Val: 0, Next: nil}, length: 0}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (ml *MyLinkedList) Get(index int) int {
    if index < 0 || index >= ml.length {
        return -1
    }
    curNode := ml.head
    for i := 0; i <= index; i++ {
        if curNode.Next == nil {
            return -1
        }
        curNode = curNode.Next
    }
    return curNode.Val
}


/** Add a node of value Val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (ml *MyLinkedList) AddAtHead(val int)  {
    node := &ListNode{Val: val, Next: nil}
    node.Next = ml.head.Next
    ml.head.Next = node
    ml.length++
}


/** Append a node of value Val to the last element of the linked list. */
func (ml *MyLinkedList) AddAtTail(val int)  {
    node := &ListNode{Val: val, Next: nil}
    curNode := ml.head
    for curNode.Next != nil {
        curNode = curNode.Next
    }
    curNode.Next = node
    ml.length++
}


/** Add a node of value Val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (ml *MyLinkedList) AddAtIndex(index int, val int)  {
    if index < 0 {
        ml.AddAtHead(val)
        return
    }
    if index == ml.length {
        ml.AddAtTail(val)
        return
    }
    if index > ml.length {
        return
    }
    preNode := ml.head
    for i := 0; i < index; i++ {
        if preNode.Next == nil {
            return
        }
        preNode = preNode.Next
    }
    node := &ListNode{Val: val, Next: nil}
    node.Next = preNode.Next
    preNode.Next = node
    ml.length++
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (ml *MyLinkedList) DeleteAtIndex(index int)  {
    if index < 0 || index >= ml.length {
        return
    }
    preNode := ml.head
    for i := 0; i < index; i++ {
        if preNode.Next == nil {
            return
        }
        preNode = preNode.Next
    }
    preNode.Next = preNode.Next.Next
    ml.length--
}

func (ml *MyLinkedList) Len() int {
    return ml.length
}

func (ml *MyLinkedList) Print() {
    fmt.Printf("-----------------\nthe list length is %d\n", ml.length)
    curNode := ml.head
    for i := 0; i < ml.length; i++ {
        if curNode.Next == nil {
            return
        }
        curNode = curNode.Next
        fmt.Printf("%d\t", curNode.Val)
    }
    fmt.Printf("\n")
}