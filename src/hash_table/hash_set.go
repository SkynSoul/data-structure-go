package hash_table

type SetNode struct {
    Val int
    Next *SetNode
    Prev *SetNode
}

type MyHashSet struct {
    data []*SetNode
}

func ConstructorHashSet() MyHashSet {
    return MyHashSet{data: make([]*SetNode, 10000)}
}

func hashFunc(key int) int {
    return key % 10000
}

func (hs *MyHashSet) Add(key int)  {
    inKey := hashFunc(key)
    inNode := &SetNode{Val: key}
    curNode := hs.data[inKey]
    if curNode == nil {
        hs.data[inKey] = inNode
        return
    }
    for {
        if curNode.Val == key {
            return
        }
        if curNode.Next == nil {
            break
        }
        curNode = curNode.Next
    }
    curNode.Next = inNode
    inNode.Prev = curNode
}

func (hs *MyHashSet) Remove(key int)  {
    inKey := hashFunc(key)
    curNode := hs.data[inKey]
    for curNode != nil {
        if curNode.Val == key {
            break
        }
        curNode = curNode.Next
    }
    if curNode == nil {
        return
    }
    if curNode.Next != nil {
        curNode.Next.Prev = curNode.Prev
    }
    if curNode.Prev == nil {
        hs.data[inKey] = curNode.Next
        return
    }
    curNode.Prev.Next = curNode.Next
}

func (hs *MyHashSet) Contains(key int) bool {
    inKey := hashFunc(key)
    curNode := hs.data[inKey]
    for curNode != nil {
        if curNode.Val == key {
            return true
        }
        curNode = curNode.Next
    }
    return false
}