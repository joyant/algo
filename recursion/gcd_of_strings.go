package recursion

/**
1071. 字符串的最大公因子
level: easy

对于字符串 s 和 t，只有在 s = t + t + t + ... + t + t（t 自身连接 1 次或多次）时，我们才认定 “t 能除尽 s”。

给定两个字符串 str1 和 str2 。返回 最长字符串 x，要求满足 x 能除尽 str1 且 x 能除尽 str2 。

示例 1：
输入：str1 = "ABCABC", str2 = "ABC"
输出："ABC"

示例 2：
输入：str1 = "ABABAB", str2 = "ABAB"
输出："AB"

示例 3：
输入：str1 = "LEET", str2 = "CODE"
输出：""
*/

func gcdOfStrings(str1 string, str2 string) string {
	// 我们总是把 str1设置成比较长的字符串，str2设置成比较短的字符串
	m, n := len(str1), len(str2)
	if m < n {
		return gcdOfStrings(str2, str1)
	}
	// 只判断较短的字符串
	if n == 0 {
		return str1
	}
	// 从较长的字符串中减去较短的，这个逻辑可以代入一个 case 看下就能明白
	if str1[:n] == str2 {
		return gcdOfStrings(str1[n:], str2)
	}
	// 其他情况返回空字符串
	return ""
}
