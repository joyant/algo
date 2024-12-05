package binary

/**
199. 二叉树的右视图
level: medium
给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
*/

// 深度优先遍历
func rightSideView(root *TreeNode) []int {
    var res []int
    rightSideViewDFS(root, &res, 0)
    return res
}

func rightSideViewDFS(root *TreeNode, res *[]int, depth int) {
    if root == nil {
        return
    }
    r := *res
    if len(r) <= depth {
        *res = append(*res, root.Val)
    } else {
        // 当前层(depth)已经被初始化了，那索引(depth)对应的值需要被重置
        // 因为遍历的顺序是中序遍历，右边的节点总是在最后被遍历到，所以最后存储的值一定是最右边的节点
        (*res)[depth] = root.Val
    }
    rightSideViewDFS(root.Left, res, depth+1)
    rightSideViewDFS(root.Right, res, depth+1)
}
