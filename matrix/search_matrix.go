package matrix

/**
面试题 10.09. 排序矩阵查找
level: medium
给定M×N矩阵，每一行、每一列都按升序排列，请编写代码找出某元素。

示例:

现有矩阵 matrix 如下：

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target = 5，返回 true。

给定 target = 20，返回 false。
*/

/*
*
三中解法：
1. 遍历所有元素，时间复杂度是 m*n
2. 遍历每个 row，因为 row 是有序的，所以可以使用二分查找，时间复杂度是 m*log n
3. z 字形遍历，从右上角开始，往左边可以找到更小的值，往下面可以找到更大的值，时间复杂度是 m+n
*/
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	i, j := 0, n-1
	for i < m && j >= 0 {
		v := matrix[i][j]
		if v == target {
			return true
		} else if v > target {
			j--
		} else {
			i++
		}
	}
	return false
}
