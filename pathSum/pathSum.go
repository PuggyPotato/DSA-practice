package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }

    // Check if it's a leaf node
    if root.Left == nil && root.Right == nil {
        return root.Val == targetSum
    }

    // Recurse on left and right with updated target
    return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

// --- Helper function to build a test tree ---
func main() {
    // Example: [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
    root := &TreeNode{Val: 5}
    root.Left = &TreeNode{Val: 4}
    root.Right = &TreeNode{Val: 8}
    root.Left.Left = &TreeNode{Val: 11}
    root.Left.Left.Left = &TreeNode{Val: 7}
    root.Left.Left.Right = &TreeNode{Val: 2}
    root.Right.Left = &TreeNode{Val: 13}
    root.Right.Right = &TreeNode{Val: 4}
    root.Right.Right.Right = &TreeNode{Val: 1}

    fmt.Println(hasPathSum(root, 22)) // true
    fmt.Println(hasPathSum(root, 26)) // true (5->8->13)
    fmt.Println(hasPathSum(root, 18)) // true (5->8->4->1)
    fmt.Println(hasPathSum(root, 5))  // false
}
