package window

/**
leetcode: 209. 长度最小的子数组

给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其总和大于等于 target 的长度最小的
子数组[numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。
如果不存在符合条件的子数组，返回 0 。

示例 1：
输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。

示例 2：
输入：target = 4, nums = [1,4,4]
输出：1

示例 3：
输入：target = 11, nums = [1,1,1,1,1,1,1,1]
输出：0
*/

// 滑动窗口
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	// i, j 分别表示 start 和 end
	i, j := 0, 0
	var result = n + 1
	var sum = 0
	for j < n {
		sum += nums[j]
		for sum >= target {
			result = min(result, j-i+1)
			// 减去 start 对应的值，左边的指针向前移动
			sum -= nums[i]
			i++
		}
		j++
	}
	if result == n+1 {
		return 0 // none
	}
	return result
}
