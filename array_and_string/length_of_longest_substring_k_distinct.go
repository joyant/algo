package array_and_string

/**
340. 至多包含 K 个不同字符的最长子串
给你一个字符串 s 和一个整数 k ，请你找出 至多 包含 k 个 不同 字符的最长
子串
，并返回该子串的长度。



示例 1：

输入：s = "eceba", k = 2
输出：3
解释：满足题目要求的子串是 "ece" ，长度为 3 。
示例 2：

输入：s = "aa", k = 1
输出：2
解释：满足题目要求的子串是 "aa" ，长度为 2 。


提示：

1 <= s.length <= 5 * 104
0 <= k <= 50
*/

// 先看懂这道题：至多包含 2 个不同字符的最长子串
// 然后再做这个就没有问题了
func lengthOfLongestSubstringKDistinct(s string, k int) int {
    n := len(s)
    if n <= k {
        return n
    }
    res := 0
    m := map[byte]int{}
    left, right := 0, 0
    for right < n {
        m[s[right]]++
        for len(m) > k {
            m[s[left]]--
            if m[s[left]] == 0 {
                delete(m, s[left])
            }
            left++
        }
        res = max(res, right-left+1)
        right++
    }
    return res
}
