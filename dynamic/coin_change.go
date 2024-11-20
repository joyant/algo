package dynamic

/**
103. 零钱兑换
给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

你可以认为每种硬币的数量是无限的。
*/

// 还是需要温习，不算完全理解
func coinChange(coins []int, amount int) int {
    max := amount + 1
    dp := make([]int, amount+1)
    for idx := range dp {
        dp[idx] = max
    }
    dp[0] = 0
    for _, coin := range coins {
        for j := coin; j <= amount; j++ {
            last := dp[j-coin]
            if last != max {
                dp[j] = min(dp[j], last+1)
            }
        }
    }
    if dp[amount] == max {
        return -1
    }
    return dp[amount]
}
