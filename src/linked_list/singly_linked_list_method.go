package linked_list

import (
    "container/list"
    "math"
)

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

func ReverseList(head *ListNode) *ListNode {
    myHead := &ListNode{Val: -1, Next: head}
    for head != nil && head.Next != nil {
        curNode := head.Next
        head.Next = curNode.Next
        curNode.Next = myHead.Next
        myHead.Next = curNode
    }
    return myHead.Next
}

func ReverseList2(head *ListNode) *ListNode {
    _, newHead := doReverseList(head)
    return newHead
}

func doReverseList(curNode *ListNode) (*ListNode, *ListNode) {
    if curNode.Next == nil {
        return curNode, curNode
    }
    preNode, newHead := doReverseList(curNode.Next)
    preNode.Next = curNode
    curNode.Next = nil
    return curNode, newHead
}

func ReverseList3(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    preNode := ReverseList(head.Next)
    head.Next.Next = head
    head.Next = nil
    return preNode
}

func RemoveElements(head *ListNode, val int) *ListNode {
    myHead := &ListNode{Val: -1, Next: head}
    preNode, curNode := myHead, myHead.Next
    for curNode != nil {
        if curNode.Val == val {
            preNode.Next = curNode.Next
        } else {
            preNode = curNode
        }
        curNode = preNode .Next
    }
    return myHead.Next
}

func OddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    evenHead := head.Next
    var oddTail, tempNode *ListNode = nil, nil
    curNode := head
    i := 1
    for curNode != nil {
        if i % 2 == 1 {
            oddTail = curNode
        }
        if curNode.Next == nil {
            break
        }
        tempNode = curNode.Next
        curNode.Next = curNode.Next.Next
        curNode = tempNode
        i++
    }
    oddTail.Next = evenHead
    return head
}

func IsPalindrome(head *ListNode) bool {
    length := 0
    curNode := head
    for curNode != nil {
        length++
        curNode = curNode.Next
    }
    halfHeadIdx := int(math.Ceil(float64(length) / 2.0))
    halfHead := head
    for i := 1; i <= halfHeadIdx; i++ {
        halfHead = halfHead.Next
    }
    halfHead = ReverseList(halfHead)
    for head != nil && halfHead != nil {
        if head.Val != halfHead.Val {
            return false
        }
        head = head.Next
        halfHead = halfHead.Next
    }
    return true
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    sentinel := &ListNode{Val: -1, Next: l1}
    var preNode, curNode, inNode, tempNode *ListNode = sentinel, l1, l2, nil
    for curNode != nil && inNode != nil {
        if inNode.Val < curNode.Val {
            tempNode = inNode.Next
            preNode.Next = inNode
            inNode.Next = curNode
            preNode = inNode
            inNode = tempNode
        } else {
            if curNode.Next == nil {
                curNode.Next = inNode
                break
            }
            preNode = curNode
            curNode = curNode.Next
        }
    }
    return sentinel.Next
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    sentinel := &ListNode{Val: -1, Next: nil}
    preNode := sentinel
    sum := 0
    for l1 != nil || l2 != nil {
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }
        preNode.Next = &ListNode{Val: sum % 10, Next: nil}
        preNode = preNode.Next
        sum = sum / 10
    }
    if sum > 0 {
        preNode.Next = &ListNode{Val: sum, Next: nil}
    }
    return sentinel.Next
}

type RandomNode struct {
    Val int
    Next *RandomNode
    Random *RandomNode
}

func CopyRandomList(head *RandomNode) *RandomNode {
    copied := make(map[*RandomNode]*RandomNode)
    return copyRandomDfs(head, copied)
}

func copyRandomDfs(head *RandomNode, copied map[*RandomNode]*RandomNode) *RandomNode{
    if head == nil {
        return head
    }
    if cpNode, ok := copied[head]; ok {
        return cpNode
    }
    curNode := &RandomNode{Val: head.Val}
    copied[head] = curNode
    curNode.Next = copyRandomDfs(head.Next, copied)
    curNode.Random = copyRandomDfs(head.Random, copied)
    return curNode
}

func CopyRandomListByStack(head *RandomNode) *RandomNode {
    if head == nil {
        return nil
    }
    stackHelper := list.New()
    copied := make(map[*RandomNode]*RandomNode)
    stackHelper.PushBack(head)
    for stackHelper.Len() > 0 {
        curNode := stackHelper.Back().Value.(*RandomNode)
        cpNode, ok := copied[curNode]
        if !ok {
            cpNode = &RandomNode{Val: curNode.Val}
            copied[curNode] = cpNode
        }
        if curNode.Next != nil {
            nextNode, ok := copied[curNode.Next]
            if ok {
                cpNode.Next = nextNode
            } else {
                stackHelper.PushBack(curNode.Next)
                continue
            }
        }
        if curNode.Random != nil {
            randomNode, ok := copied[curNode.Random]
            if ok {
                cpNode.Random = randomNode
            } else {
                stackHelper.PushBack(curNode.Random)
                continue
            }
        }
        stackHelper.Remove(stackHelper.Back())
    }
    return copied[head]
}