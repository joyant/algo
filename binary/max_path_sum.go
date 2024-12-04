package binary

import "math"

/**
124. 二叉树中的最大路径和
level: hard
二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。

路径和 是路径中各节点值的总和。

给你一个二叉树的根节点 root ，返回其 最大路径和 。
*/

func maxPathSum(root *TreeNode) int {
    maxSum := math.MinInt32
    var maxGain func(node *TreeNode) int
    maxGain = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        // 计算左边分支最大贡献值，只有在贡献值大于 0 时，才会选取对应分支
        leftGain := max(maxGain(node.Left), 0)
        // 计算右边分支最大贡献值，只有在贡献值大于 0 时，才会选取对应分支
        rightGain := max(maxGain(node.Right), 0)

        // 当前节点的值和通过当前节点的路径和作为候选答案
        priceNewpath := node.Val + leftGain + rightGain

        // 更新答案
        maxSum = max(maxSum, priceNewpath)

        // 返回节点的最大贡献值
        return node.Val + max(leftGain, rightGain)
    }
    maxGain(root)
    return maxSum
}
