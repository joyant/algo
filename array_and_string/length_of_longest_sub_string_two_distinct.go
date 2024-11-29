package array_and_string

/**
159. 至多包含两个不同字符的最长子串

给你一个字符串 s ，请你找出 至多 包含 两个不同字符 的最长
子串
，并返回该子串的长度。


示例 1：

输入：s = "eceba"
输出：3
解释：满足题目要求的子串是 "ece" ，长度为 3 。
示例 2：

输入：s = "ccaabbb"
输出：5
解释：满足题目要求的子串是 "aabbb" ，长度为 5 。


提示：

1 <= s.length <= 105
s 由英文字母组成
*/

// 双指针
// 使用一个 map，中间删除动作是一个奇妙的操作，令逻辑大大简化
func lengthOfLongestSubstringTwoDistinct(s string) int {
    n := len(s)
    if n <= 2 {
        return n
    }
    res := 0
    m := map[byte]int{}
    left, right := 0, 0
    for right < n {
        r := s[right]
        m[r]++
        for len(m) > 2 {
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
