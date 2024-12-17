package binary

/**
637. 二叉树的层平均值
level: easy
给定一个非空二叉树的根节点 root , 以数组的形式返回每一层节点的平均值。与实际答案相差 10-5 以内的答案可以被接受。
*/

type levelData struct {
    sum, count int
}

func averageOfLevels(root *TreeNode) []float64 {
    var levelValues []levelData
    var dfs func(*TreeNode, int)
    dfs = func(root *TreeNode, level int) {
        if root == nil {
            return
        }
        if level < len(levelValues) {
            levelValues[level].sum += root.Val
            levelValues[level].count++
        } else {
            levelValues = append(levelValues, levelData{sum: root.Val, count: 1})
        }
        dfs(root.Left, level+1)
        dfs(root.Right, level+1)
    }
    dfs(root, 0)
    res := make([]float64, 0, len(levelValues))
    for _, v := range levelValues {
        res = append(res, float64(v.sum)/float64(v.count))
    }
    return res
}
