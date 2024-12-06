package array_and_string

/**
1180. 统计只含单一字母的子串
level:easy
给你一个字符串 s，返回 只含 单一字母 的子串个数 。

示例 1：
输入： s = "aaaba"
输出： 8
解释： 只含单一字母的子串分别是 "aaa"， "aa"， "a"， "b"。
"aaa" 出现 1 次。
"aa" 出现 2 次。
"a" 出现 4 次。
"b" 出现 1 次。
所以答案是 1 + 2 + 4 + 1 = 8。

示例 2:
输入： s = "aaaaaaaaaa"
输出： 55
*/

/*
*
有两个地方需要注意
1. 读懂题干
这个题目的意思是，给你一个字符串，你需要计算出这个字符串中所有只包含单一字母的子串的总数。所谓"只包含单一字母的子串"，就是这个子串中的字母都是一样的。

举个例子，假设给定字符串 "aaaba"。这个字符串中，只包含单一字母的子串有 "aaa"、"aa"、"a" 和 "b"。
"aaa" 出现了 1 次，"aa" 出现了 2 次，"a" 出现了 4 次，"b" 出现了 1 次。所以总的子串数就是这些出现次数的总和，即 1 + 2 + 4 + 1 = 8。

再比如，如果给定字符串 "aaaaaaaaaa"，那么只包含单一字母的子串就有 "aaaaaaaaaa"、"aaaaaaaaa"、"aaaaaaaa" 等等，
每个子串都只包含字母 "a"。这些子串的出现次数分别是 1、2、3、...、10，所以总的子串数就是这些出现次数的总和，即 1 + 2 + 3 + ... + 10 = 55。

所以，你需要编写一个函数，输入一个字符串，输出这个字符串中所有只包含单一字母的子串的总数。
2. 通用的公式是：n * (n + 1) / 2，比如 aaa 共有 6 个子串。a 3 个， aa 2 个，aaa 1 个
*/
func countLetters(s string) int {
	var prev byte
	res, count := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] != prev {
			res += calcCount(count)
			count = 1
			prev = s[i]
		} else {
			count++
		}
	}
	res += calcCount(count)
	return res
}

func calcCount(c int) int {
	return c * (c + 1) / 2
}
