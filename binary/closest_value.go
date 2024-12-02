package binary

/**
270. 最接近的二叉搜索树值
level: easy
给你二叉搜索树的根节点 root 和一个目标值 target ，请在该二叉搜索树中找到最接近目标值 target 的数值。如果有多个答案，返回最小的那个。
*/

/**
 二叉搜索树的性质是，对于每个节点，其左子树中的所有元素的值都小于该节点的值，而其右子树中的所有元素的值都大于该节点的值。

基于这个性质，我们可以从根节点开始，比较当前节点的值和目标值的大小，如果当前节点的值大于目标值，那么我们就往左子树走，否则就往右子树走。在这个过程中，我们还需要一个变量来记录我们遇到的与目标值最接近的节点的值。

个人觉得这个题把 target 设置成 int 也没有问题，非要设置成 float64
需要加一些必要的类型转换，整个算法看起来很冗繁
*/

func closestValue(root *TreeNode, target float64) int {
    res := root.Val
    // 不必要判断停止了，dfs 到最后
    for root != nil {
        cur := floatAbs(float64(root.Val) - target)
        prev := floatAbs(float64(res) - target)
        // 注意题干的一句话：如果有多个答案，返回最小的那个。
        if cur == prev {
            res = min(res, root.Val)
        } else if cur < prev {
            res = root.Val
        }
        if target < float64(root.Val) {
            root = root.Left
        } else {
            root = root.Right
        }
    }
    return res
}

func floatAbs(f float64) float64 {
    if f < 0 {
        return -f
    }
    return f
}
