package window

/**
给你一个字符串 S，找出所有长度为 K 且不含重复字符的子串，请你返回全部满足要求的子串的 数目。
level: medium
示例 1：

输入：S = "havefunonleetcode", K = 5
输出：6
解释：
这里有 6 个满足题意的子串，分别是：'havef','avefu','vefun','efuno','etcod','tcode'。
示例 2：

输入：S = "home", K = 5
输出：0
解释：
注意：K 可能会大于 S 的长度。在这种情况下，就无法找到任何长度为 K 的子串。


提示：

1 <= S.length <= 10^4
S 中的所有字符均为小写英文字母
1 <= K <= 10^4
*/

// 滑动窗口
func numKLenSubstrNoRepeats(s string, k int) int {
    n := len(s)
    if n < k {
        return 0
    }
    count := [26]int{}
    left, right, res := 0, 0, 0
    for right < n {
        count[s[right]-'a']++
        // 关键是这里的 for 循环的条件
        for right-left+1 > k || count[s[right]-'a'] > 1 {
            count[s[left]-'a']--
            left++
        }
        if right-left+1 == k {
            res++
        }
        right++
    }
    return res
}
