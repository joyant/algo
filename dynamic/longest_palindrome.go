package dynamic

/**
5. 最长回文子串
level: medium
给你一个字符串 s，找到 s 中最长的
回文子串。

示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"
*/

/**
这个问题可以通过动态规划来解决。我们可以创建一个二维数组 dp，其中 dp[i][j] 表示字符串 s 的第 i 个字符到第 j 个字符组成的子串是否是回文串。

如果一个字符串的头尾两个字符都不相等，那么这个字符串一定不是回文串；如果一个字符串的头尾两个字符相等，这里就要分情况讨论了：

如果这个字符串长度不大于 3，那么它就是一个回文串；
如果这个字符串长度大于 3，那么其是否是回文串取决于 dp[i + 1][j - 1] 是否是回文串，也就是说，只有当它的子串是回文串，并且头尾两个字符相等时，它才是回文串。

例子：
      a     b     b     a
    +-----+-----+-----+-----+
  a |  T  |  F  |  F  |  T  |
    +-----+-----+-----+-----+
  b |     |  T  |  T  |  F  |
    +-----+-----+-----+-----+
  b |     |     |  T  |  F  |
    +-----+-----+-----+-----+
  a |     |     |     |  T  |
    +-----+-----+-----+-----+
*/

func longestPalindrome(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}
	dp := make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
		dp[i][i] = true
	}
	var maxLen = 1
	var start = 0
	for j := 1; j < length; j++ {
		// i < j 就可以，没有必要 i < length
		for i := 0; i < length; i++ {
			if s[i] == s[j] {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			} else {
				dp[i][j] = false
			}
			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				start = i
			}
		}
	}
	printMatrix(dp)
	return s[start : start+maxLen]
}
