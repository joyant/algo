package tree

/**
226. 翻转二叉树

给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
*/

// 这是我自己写的，官方的更简洁，但我自己写的更容易懂
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    left := root.Left
    right := root.Right
    root.Left = right
    root.Right = left
    invertTree(left)
    invertTree(right)
    return root
}
