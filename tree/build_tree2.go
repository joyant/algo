package tree

/**
106. 从中序与后序遍历序列构造二叉树
给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并返回这颗 二叉树 。
*/

func buildTree2(inorder []int, postorder []int) *TreeNode {
    if len(inorder) == 0 {
        return nil
    }
    // 由postorder 可以获得根节点
    root := &TreeNode{Val: postorder[len(postorder)-1]}
    // 找到 root 节点的 index
    i := 0
    for ; i < len(inorder); i++ {
        if inorder[i] == root.Val {
            break
        }
    }
    root.Left = buildTree2(inorder[:i], postorder[:i])
    root.Right = buildTree2(inorder[i+1:], postorder[i:len(postorder)-1])
    return root
}
