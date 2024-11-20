package tree

/**
LCR 144. 翻转二叉树
给定一棵二叉树的根节点 root，请左右翻转这棵二叉树，并返回其根节点。
*/

// 递归
func flipTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    root.Left, root.Right = root.Right, root.Left
    flipTree(root.Left)
    flipTree(root.Right)
    return root
}
