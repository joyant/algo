package array_and_string

/**
162. 寻找峰值
level: medium

峰值元素是指其值严格大于左右相邻值的元素。

给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。

你可以假设 nums[-1] = nums[n] = -∞ 。

你必须实现时间复杂度为 O(log n) 的算法来解决此问题。
*/

// level：从代码看是 easy，从思路看是 medium
// 可以用[1, 3, 2]作为例子，只要3比1大，且没有找到比3更大的值，那么3就是我们要的答案
// 没有特有名词可以概括，就叫遍历吧
// 感觉这样的解法的前提是，给出的 nums 数组中必须有答案
func findPeakElement(nums []int) int {
	res := 0
	for idx, num := range nums {
		if num > nums[res] {
			res = idx
		}
	}
	return res
}
