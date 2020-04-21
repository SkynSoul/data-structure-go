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
    if index >= ml.length {
        ml.AddAtTail(val)
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

func HasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil || head.Next.Next == nil {
        return false
    }
    slow, fast := head.Next, head.Next.Next
    for slow != fast {
        if fast == nil || fast.Next == nil {
            return false
        }
        fast = fast.Next.Next
        if slow == nil {
            return false
        }
        slow = slow.Next
    }
    return fast != nil
}

func DetectCycle(head *ListNode) *ListNode {
    nodeMap := make(map[*ListNode]bool)
    curNode := head
    for curNode != nil {
        _, ok := nodeMap[curNode]
        if ok {
            return curNode
        }
        nodeMap[curNode] = true
        curNode = curNode.Next
    }
    return nil
}

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
    for headA != nil {
        curB := headB
        for curB != nil {
            if headA == curB {
                return headA
            }
            curB = curB.Next
        }
        headA = headA.Next
    }
    return nil
}

// 设交点之前A的路径长度为a，B的路径长度为b，公共的为n，默认a > b
// B走到终点，A还剩a - b
// B从头走A的路径
// A走到终点，B走了a-b
// A从头走B的路径
// A走到相交点时，走了b的长度，同时B也走了a - b + b = a的长度，恰好也在相交点
func GetIntersectionNode2(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }
    curA, curB := headA, headB
    var tailA, tailB *ListNode = nil, nil
    for {
        if curA == curB {
            return curA
        }
        if curA.Next != nil {
            curA = curA.Next
        } else if tailA == nil {
            tailA = curA
            curA = headB
        }
        if curB.Next != nil {
            curB = curB.Next
        } else if tailB == nil {
            tailB = curB
            curB = headA
        }
        if tailA != nil && tailB != nil && tailA != tailB {
            return nil
        }
    }
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
    arrHelper := make([]*ListNode, n + 1)
    i := -1
    curNode := head
    for curNode != nil {
        i++
        arrHelper[i % (n + 1)] = curNode
        curNode = curNode.Next
    }
    delNode := arrHelper[(i + 2) % (n + 1)]
    if delNode == head {
        return head.Next
    }
    preNode := arrHelper[(i + 1) % (n + 1)]
    preNode.Next = preNode.Next.Next
    return head
}