package array_and_string

/**
给定 m 个数组，每个数组都已经按照升序排好序了。
level: medium (maybe easy)
现在你需要从两个不同的数组中选择两个整数（每个数组选一个）并且计算它们的距离。两个整数 a 和 b 之间的距离定义为它们差的绝对值 |a-b| 。
返回最大距离。

示例 1：
输入：[[1,2,3],[4,5],[1,2,3]]
输出：4
解释：
一种得到答案 4 的方法是从第一个数组或者第三个数组中选择 1，同时从第二个数组中选择 5 。

示例 2：
输入：arrays = [[1],[1]]
输出：0
*/

// 中等难度，实际上应该是简单
func maxDistance(arrays [][]int) int {
	n := len(arrays)
	if n == 0 {
		return 0
	}
	minVal, maxVal := arrays[0][0], arrays[0][len(arrays[0])-1]
	result := 0
	for i := 1; i < n; i++ {
		curMin, curMax := arrays[i][0], arrays[i][len(arrays[i])-1]
		result = max(result, abs(curMin-maxVal), abs(curMax-minVal))
		minVal = min(minVal, curMin)
		maxVal = max(maxVal, curMax)
	}
	return result
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}
