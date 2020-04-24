package hash_table

type MapNode struct {
    Key     int
    Val     int
    Next    *MapNode
}

type MapBucket struct {
    head    *MapNode
}

func (mb *MapBucket) PutPair(key int, val int) {
    curNode := mb.head
    for curNode.Next != nil {
        if curNode.Next.Key == key {
            curNode.Next.Val = val
            return
        }
        curNode = curNode.Next
    }
    curNode.Next = &MapNode{Key: key, Val: val}
}

func (mb *MapBucket) GetPair(key int) *MapNode {
    curNode := mb.head.Next
    for curNode != nil {
        if curNode.Key == key {
            break
        }
        curNode = curNode.Next
    }
    return curNode
}

func (mb *MapBucket) RemovePair(key int) {
    curNode := mb.head
    for curNode.Next != nil {
        if curNode.Next.Key == key {
            curNode.Next = curNode.Next.Next
            return
        }
        curNode = curNode.Next
    }
}

type MyHashMap struct {
    data        []*MapBucket
    keyRange    int
}

func ConstructorHashMap() MyHashMap {
    return MyHashMap{data: make([]*MapBucket, 6151), keyRange: 6151}
}

func (hm *MyHashMap) hashFunc(key int) int {
    return key % hm.keyRange
}

func (hm *MyHashMap) Put(key int, value int) {
    inKey := hm.hashFunc(key)
    if hm.data[inKey] == nil {
        hm.data[inKey] = &MapBucket{head: &MapNode{Key: 0, Val: 0}}
    }
    hm.data[inKey].PutPair(key, value)
}

func (hm *MyHashMap) Get(key int) int {
    inKey := hm.hashFunc(key)
    if hm.data[inKey] == nil {
        return -1
    }
    retNode := hm.data[inKey].GetPair(key)
    if retNode == nil {
        return -1
    }
    return retNode.Val
}

func (hm *MyHashMap) Remove(key int) {
    inKey := hm.hashFunc(key)
    if hm.data[inKey] == nil {
        return
    }
    hm.data[inKey].RemovePair(key)
}