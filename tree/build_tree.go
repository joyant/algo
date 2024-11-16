package tree

/**
105. 从前序与中序遍历序列构造二叉树
给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
*/

func buildTree(preorder []int, inorder []int) *TreeNode {
    // 这里判断 inorder 也行，效果一样
    if len(preorder) == 0 {
        return nil
    }
    // 前序遍历的第一个元素一定是 root 节点
    root := &TreeNode{Val: preorder[0]}
    // 从中序遍历的结果中找到 root 节点的 index
    // 不知道这个题为什么没有加「元素节点的值不会重复」的条件
    i := 0
    for ; i < len(inorder); i++ {
        if inorder[i] == root.Val {
            break
        }
    }
    // 可以利用 i，它实际上是一个长度值
    // inorder[:i]是中序遍历的左节点中序的遍历结果, [i+1:]是右节点中序遍历结果
    // preorder 的 1:i+1 的元素是左节点前序遍历的结果，[i+1:]是右节点前序遍历的结果
    root.Left = buildTree(preorder[1:i+1], inorder[:i])
    root.Right = buildTree(preorder[i+1:], inorder[i+1:])
    return root
}
