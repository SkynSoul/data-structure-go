package stack

type Node struct {
	Val int
	Neighbors []*Node
}

func CloneGraph1(node *Node) *Node {
	start := &Node{Val: node.Val, Neighbors: make([]*Node, 0)}
	visited := make(map[int]*Node)
	dfsGraph1(node, start, visited)
	return start
}

func dfsGraph1(src *Node, dst *Node, visited map[int]*Node) {
	if src == nil {
		return
	}
	_, ok := visited[src.Val]
	if ok {
		return
	}
	visited[dst.Val] = dst
	for i := 0; i < len(src.Neighbors); i++ {
		next, ok := visited[src.Neighbors[i].Val]
		if !ok {
			next = &Node{Val: src.Neighbors[i].Val, Neighbors: make([]*Node, 0)}
		}
		dst.Neighbors = append(dst.Neighbors, next)
		dfsGraph1(src.Neighbors[i], next, visited)
	}
}

// 第二种方式比较优雅
func CloneGraph2(node *Node) *Node {
	visited := make(map[int]*Node)
	return dfsGraph2(node, visited)
}

func dfsGraph2(src *Node, visited map[int]*Node) *Node {
	if src == nil {
		return nil
	}
	curNode, ok := visited[src.Val]
	if ok {
		return curNode
	}
	dst := &Node{Val: src.Val, Neighbors: make([]*Node, 0)}
	visited[dst.Val] = dst
	for i := 0; i < len(src.Neighbors); i++ {
		dst.Neighbors = append(dst.Neighbors, dfsGraph2(src.Neighbors[i], visited))
	}
	return dst
}