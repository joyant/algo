package tree

/**
101. 对称二叉树
给你一个二叉树的根节点 root ， 检查它是否轴对称。
*/

// 递归
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return check(root.Left, root.Right)
}

func check(a, b *TreeNode) bool {
    if a == nil && b == nil {
        return true
    }
    if a == nil || b == nil {
        return false
    }
    return a.Val == b.Val && check(a.Left, b.Right) && check(a.Right, b.Left)
}
