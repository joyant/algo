package tree

/**
222. 完全二叉树的节点个数
给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。

完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，
并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。
*/

func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    lDepth := depth(root.Left)
    rDepth := depth(root.Right)
    //  第一次执行 countNodes 方法，因为左右的深度相等
    // 所以判断左边是完全二叉树，左边的节点数一定是 2^lDepth
    // 然后再计算右子树
    if lDepth == rDepth {
        return (1 << lDepth) + countNodes(root.Right)
    } else {
        return (1 << rDepth) + countNodes(root.Left)
    }
}

func depth(root *TreeNode) int {
    d := 0
    for root != nil {
        d++
        root = root.Left
    }
    return d
}
