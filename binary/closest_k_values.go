package binary

// 用一个数组，保存前 k 个值，前序遍历的情况下，能保证数组中的值是有序递增的
// 这里的难点是为什么每次root.Val 都和 res[0]做比较
// 举个例子，如果现在 res 的值是[1, 2]，k 的值是2，root.Val 一定是比2大，而不是比2小
// 比1和2都大且与1相比更接近target 的情况下，2肯定比1更接近
// 这就是为什么总是 remove 掉 res[0]，然后把 root.Val append 到末尾的原因
func closestKValues(root *TreeNode, target float64, k int) []int {
    res := make([]int, 0, k)
    dfs(root, target, k, &res)
    return res
}

func dfs(root *TreeNode, target float64, k int, res *[]int) {
    if root == nil {
        return
    }
    dfs(root.Left, target, k, res)
    if len(*res) < k {
        *res = append(*res, root.Val)
    } else {
        if abs(float64((*res)[0])-target) > abs(float64(root.Val)-target) {
            *res = append((*res)[:0], (*res)[1:]...)
            *res = append(*res, root.Val)
        }
    }
    dfs(root.Right, target, k, res)
}

func abs(n float64) float64 {
    if n > 0 {
        return n
    }
    return -n
}
