package stack

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// 递归
func InOrderTraversal(root *TreeNode) []int {
    ret := make([]int, 0)
    visited := make(map[*TreeNode]bool)
    recursionTraversal(root, visited, &ret)
    return ret
}

func recursionTraversal(root *TreeNode, visited map[*TreeNode]bool, ret *[]int)  {
    if root == nil {
        return
    }
    _, ok := visited[root]
    if ok {
        return
    }
    visited[root] = true
    recursionTraversal(root.Left, visited, ret)
    *ret = append(*ret, root.Val)
    recursionTraversal(root.Right, visited, ret)
}

// 辅助栈
func InOrderTraversal2(root *TreeNode) []int {
    ret := make([]int, 0)
    if root == nil {
       return ret
    }
    visited := make(map[*TreeNode]bool)
    stackHelper := make([]*TreeNode, 0)
    stackHelper = append(stackHelper, root)
    visited[root] = true
    for len(stackHelper) > 0 {
        curNode := stackHelper[len(stackHelper) - 1]
        if curNode.Left != nil {
            if _, ok := visited[curNode.Left]; !ok {
                stackHelper = append(stackHelper, curNode.Left)
                visited[curNode.Left] = true
                continue
            }
        }
        ret = append(ret, curNode.Val)
        stackHelper = stackHelper[:len(stackHelper) - 1]
        if curNode.Right != nil {
            if _, ok := visited[curNode.Right]; !ok {
                stackHelper = append(stackHelper, curNode.Right)
                visited[curNode.Right] = true
            }
        }
    }
    return ret
}