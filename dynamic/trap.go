package dynamic

/**
42. 接雨水
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
*/
/**
java 版本：
public int trap(int[] height) {
    int sum = 0;
    int[] max_left = new int[height.length];
    int[] max_right = new int[height.length];

    for (int i = 1; i < height.length - 1; i++) {
        max_left[i] = Math.max(max_left[i - 1], height[i - 1]);
    }
    for (int i = height.length - 2; i >= 0; i--) {
        max_right[i] = Math.max(max_right[i + 1], height[i + 1]);
    }
    for (int i = 1; i < height.length - 1; i++) {
        int min = Math.min(max_left[i], max_right[i]);
        if (min > height[i]) {
            sum = sum + (min - height[i]);
        }
    }
    return sum;
}

作者：windliang
链接：https://leetcode.cn/problems/trapping-rain-water/description/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func trap(height []int) int {
    length := len(height)
    maxLeft := make([]int, length)
    maxRight := make([]int, length)
    // fill max left
    for i := 1; i < length; i++ {
        maxLeft[i] = max(maxLeft[i-1], height[i-1])
    }
    // fill max right
    for i := length - 2; i >= 0; i-- {
        maxRight[i] = max(maxRight[i+1], height[i+1])
    }
    var sum int
    // calc max
    for i := 1; i < length-1; i++ {
        min := min(maxLeft[i], maxRight[i])
        if min > height[i] {
            sum += (min - height[i])
        }
    }
    return sum
}
