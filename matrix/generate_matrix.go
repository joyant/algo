package matrix

/**
59. 螺旋矩阵 II
level: medium
给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
*/

// 如何用笔自然的在纸上画出矩阵，就如何按顺序填充，多做就熟悉了，没有技术含量，纯粹炫技
func generateMatrix(n int) [][]int {
	// 初始化 matrix
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	counter := 1
	// 填充，按 counter 递增的方向填充
	left, right, top, bottom := 0, n-1, 0, n-1
	// left right 对应 j，是横向
	// top bottom 对应 i，是竖向
	for left <= right && top <= bottom {
		// 填充上边
		for idx := left; idx <= right; idx++ {
			matrix[top][idx] = counter
			counter++
		}
		top++
		// 填充右边
		for idx := top; idx <= bottom; idx++ {
			matrix[idx][right] = counter
			counter++
		}
		right--
		// 填充下边
		for idx := right; idx >= left; idx-- {
			matrix[bottom][idx] = counter
			counter++
		}
		bottom--
		// 填充左边
		for idx := bottom; idx >= top; idx-- {
			matrix[idx][left] = counter
			counter++
		}
		left++
	}
	return matrix
}
