package array_and_string

/**
leetcode:11 盛最多水的容器

给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

说明：你不能倾斜容器。
*/

// 这是一个双指针问题
// 本质上是求两边的墙组成和矩形面积
func maxArea(height []int) int {
    n := len(height)
    i, j := 0, n-1
    var result int
    for i < j {
        result = max(result, min(height[i], height[j])*(j-i))
        // 尽量让指针在高的墙这里停留时间长一点
        if height[i] < height[j] {
            i++
        } else {
            j--
        }
    }
    return result
}
