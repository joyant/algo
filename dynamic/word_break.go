package dynamic

/**
leetcode: 139
给你一个字符串 s 和一个字符串列表 wordDict 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true。

注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。

示例 1：

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
示例 2：

输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
     注意，你可以重复使用字典中的单词。
示例 3：

输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false

这是一个很好的问题，让我来解释一下：

1. **为什么长度是 n+1 而不是 n？**

   我们的 dp 数组是用来表示字符串 s 的前 i 个字符能否用 wordDict 中的单词表示。这里的 "前 i 个字符" 是包含 0 的，也就是说，当 i=0 时，我们实际上是在表示空字符串。所以，我们需要 n+1 个元素来表示从空字符串到整个字符串的所有可能。

2. **为什么 dp[0] 是 true？**

   dp[0] 实际上表示的是空字符串 "" 能否被 wordDict 中的单词表示。因为空字符串不需要任何单词就可以表示（也就是说，我们不使用任何单词就可以得到空字符串），所以我们将 dp[0] 初始化为 true。

这种方式在动态规划问题中很常见，我们经常需要额外的空间来处理边界情况，而且我们经常会把这些边界情况初始化为某种我们知道的值，以便我们可以从这些已知的值开始填充我们的 dp 数组。

通常情况下，我们更倾向于使用 n+1 长度的 dp 数组并将 dp[0] 设置为 true，而不是特殊处理空字符串的情况，这样可以使代码更加统一和简洁。这种方式也更符合动态规划的一般模式，即从基础情况开始并基于这些基础情况构建更复杂的情况。
*/
func wordBreak(s string, wordDict []string) bool {
    n := len(s)
    set := make(map[string]bool)
    for _, w := range wordDict {
        set[w] = true
    }
    dp := make([]bool, n+1)
    dp[0] = true // 空字符串的话，函数的结果应该是 true
    // 为什么是 i<=n，因为并不会使用 n
    for i := 1; i <= n; i++ {
        // 为什么时 j<n，因为考虑下 s 只是一个字符的情况
        for j := 0; j < i; j++ {
            // j 讲当前 s 分为两段，每段都是一个完整的单词
            if dp[j] && set[s[j:i]] {
                dp[i] = true
                break
            }
        }
    }
    return dp[n]
}
