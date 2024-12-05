package binary

/**
199. 二叉树的右视图
level: medium
给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
*/

/**
这个问题是关于二叉树的，要求我们找到从二叉树的右侧看过去的视图。换句话说，对于每一层，我们只需要找到最右边的节点。

这个解决方案采用了深度优先搜索（DFS）的方法。它首先创建一个空的列表 `ans` 来存储每一层最右边的节点。然后，它调用 `dfs` 函数，以根节点 `root` 和深度 `0` 作为参数。

在 `dfs` 函数中，它首先检查当前节点 `node` 是否为空。如果为空，则返回。这是递归的基本情况。

然后，它检查 `ans` 的长度是否小于或等于当前的深度。如果是，那么这是我们在这一层遇到的第一个节点，我们将它的值添加到 `ans` 中。否则，我们已经在这一层添加过一个节点，我们需要更新它的值，因为我们想要的是这一层最右边的节点。

最后，我们对左子节点和右子节点进行递归调用，深度增加 1。

这个解决方案的关键在于，它首先访问左子节点，然后访问右子节点。这样，对于每一层，最后一个被访问的节点就是最右边的节点，这就是我们想要的结果。

这个解决方案的时间复杂度是 O(n)，其中 n 是二叉树的节点数，因为我们需要访问每一个节点。空间复杂度也是 O(n)，因为在最坏的情况下，我们可能需要存储所有节点的值。
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
