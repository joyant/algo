package tree

/**
114. 二叉树展开为链表
给你二叉树的根结点 root ，请你将它展开为一个单链表：

展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。
*/

func flatten(root *TreeNode) {
    if root == nil {
        return
    }
    // 先把子节点都 flatten 了，然后操作当前节点的左右节点
    flatten(root.Left)
    flatten(root.Right)
    left, right := root.Left, root.Right
    root.Left = nil
    root.Right = left
    // 找到移动过来的 right 的最后一个节点，指针现在的 right
    p := root
    for p.Right != nil {
        p = p.Right
    }
    p.Right = right
}
