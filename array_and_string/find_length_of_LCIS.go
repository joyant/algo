package array_and_string

/**
674. 最长连续递增序列
level: easy
给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。

连续递增的子序列 可以由两个下标 l 和 r（l < r）确定，如果对于每个 l <= i < r，都有 nums[i] < nums[i + 1] ，
那么子序列 [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] 就是连续递增子序列。

示例 1：
输入：nums = [1,3,5,4,7]
输出：3
解释：最长连续递增序列是 [1,3,5], 长度为3。
尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为 5 和 7 在原数组里被 4 隔开。

示例 2：
输入：nums = [2,2,2,2,2]
输出：1
解释：最长连续递增序列是 [2], 长度为1。
*/

// 用两个指针i 和 j，i 不动，向右移动 j，如果不在递增，重置 i 和 j 的值
func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	res := 1
	i, j := 0, 1
	for j < n {
		if nums[j] > nums[j-1] {
			res = max(res, j-i+1)
			j++
		} else {
			i = j
			j = i + 1
		}
	}
	return res
}
