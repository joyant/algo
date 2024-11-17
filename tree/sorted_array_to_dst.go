package tree

/**
108. 将有序数组转换为二叉搜索树
给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵
平衡
 二叉搜索树。
*/

// 一个典型的递归算法，但又是一个短时间内，想不到的话就不会做的算法
// 数组的中间节点一定是树的根节点，把 nums 从中间分开后，再在左右两个数组上应用前面我说过的话
func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    mid := len(nums) / 2
    left := sortedArrayToBST(nums[:mid])
    right := sortedArrayToBST(nums[mid+1:])
    return &TreeNode{Val: nums[mid], Left: left, Right: right}
}
