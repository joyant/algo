package array_and_string

/**
leetcode: 3. 无重复字符的最长子串

给定一个字符串 s ，请你找出其中不含有重复字符的 最长
子串的长度。

示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

// 与 209 相似，也是滑动窗口，但这里需要利用 hash
func lengthOfLongestSubstring(s string) int {
    start, maxLen := 0, 0
    m := make(map[rune]int)
    for end, char := range s {
        // 这一步真的很妙，虽然可能有很多字符重复，但保留的值永远是最后一次出现的 index
        if lastSeenPos, ok := m[char]; ok && lastSeenPos >= start {
            start = lastSeenPos + 1
        }
        maxLen = max(maxLen, end-start+1)
        // 点睛
        m[char] = end
    }
    return maxLen
}
