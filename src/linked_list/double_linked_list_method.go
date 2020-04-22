package linked_list

type Node struct {
    Val int
    Prev *Node
    Next *Node
    Child *Node
}

func Flatten(root *Node) *Node {
    if root == nil {
        return nil
    }
    root, _ = doFlatten(root)
    return root
}

func doFlatten(root *Node) (*Node, *Node)  {
    if root.Child == nil && root.Next == nil {
        return root, root
    }
    curHead, curTail, rootNext := root, root, root.Next
    if root.Child != nil {
        childHead, childTail := doFlatten(root.Child)
        curHead.Next = childHead
        childHead.Prev = curHead
        curTail = childTail
        root.Child = nil
    }
    if rootNext != nil {
        nextHead, nextTail := doFlatten(rootNext)
        curTail.Next = nextHead
        nextHead.Prev = curTail
        curTail = nextTail
    }
    return curHead, curTail
}