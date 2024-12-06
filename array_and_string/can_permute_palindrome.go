package array_and_string

/**
266. 回文排列

给你一个字符串 s ，如果该字符串的某个排列是
回文串 ，则返回 true ；否则，返回 false 。

示例 1：
输入：s = "code"
输出：false

示例 2：
输入：s = "aab"
输出：true

示例 3：
输入：s = "carerac"
输出：true
*/

// 理解题干很重要，「某个排列」的意思是，你可以重新的排列组合 s，只要能组成回文，那 s 就是回文
// 所以可能性就两种：
// 1. 所有字符都出现偶数次
// 2. 只有一个字符出现奇数次，其他字符都是偶数次
// 这样的话，适当的排列组合后，一定能产生回文

func canPermutePalindrome(s string) bool {
	mp := make(map[rune]int)
	for _, c := range s {
		mp[c]++
	}
	count := 0
	for _, v := range mp {
		count += v % 2
	}
	return count <= 1
}
