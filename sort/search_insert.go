package sort

/**
35. 搜索插入位置
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 O(log n) 的算法。
*/

// 二分查找
func searchInsert(nums []int, target int) int {
    if len(nums) == 0 {
        return 0
    }
    left, right := 0, len(nums)-1
    for left <= right {
        mid := (right-left)/2 + left
        v := nums[mid]
        if v == target {
            return mid
        } else if v < target { // 这里为什么不是<=?
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return left
}
