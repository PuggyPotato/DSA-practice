package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
    // If both nodes are nil, they are the same
    if p == nil && q == nil {
        return true
    }
    // If one is nil and the other isn't, they are not the same
    if p == nil || q == nil {
        return false
    }
    // Check current node value and recursively check left and right subtrees
    return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// Helper function to build tree (optional, for testing)
func main() {
    p := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
    q := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}

    fmt.Println(isSameTree(p, q)) // Output: true

    r := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
    s := &TreeNode{1, nil, &TreeNode{2, nil, nil}}

    fmt.Println(isSameTree(r, s)) // Output: false
}
