package greedy

/**
45. 跳跃游戏 II
level: medium

给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。

每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:

0 <= j <= nums[i]
i + j < n
返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。
*/

/**
解题思路
贪心思想：

在每次跳跃时，尽量选择能跳得最远的地方，保证跳跃次数最少。
我们维护一个变量表示当前能够到达的最远位置，并通过遍历数组来更新这个值。
关键变量：

jumps：记录跳跃次数。
currentEnd：表示当前跳跃的覆盖范围。
farthest：记录能到达的最远位置。
遍历数组：

从左到右遍历数组，更新 farthest。
当到达 currentEnd 时，增加跳跃次数并更新 currentEnd 为 farthest。
终止条件：

一旦 currentEnd 覆盖到数组末尾即可返回 jumps。
*/

func jump(nums []int) int {
    n := len(nums)
    // 元素长度为0或者1表示不用跳，没有元素无法跳，只有一个，起点就是终点，不用跳
    if n <= 1 {
        return 0
    }
    jumps, currentEnd, farthest := 0, 0, 0
    for i := 0; i < n; i++ {
        farthest = max(farthest, i+nums[i])
        if i == currentEnd {
            jumps++
            currentEnd = farthest
            if currentEnd >= n-1 {
                break
            }
        }
    }
    return jumps
}
