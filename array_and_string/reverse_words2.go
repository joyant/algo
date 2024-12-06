package array_and_string

/**
186. 反转字符串中的单词 II

给你一个字符数组 s ，反转其中 单词 的顺序。

单词 的定义为：单词是一个由非空格字符组成的序列。s 中的单词将会由单个空格分隔。

必须设计并实现 原地 解法来解决此问题，即不分配额外的空间。

示例 1：
输入：s = ["t","h","e"," ","s","k","y"," ","i","s"," ","b","l","u","e"]
输出：["b","l","u","e"," ","i","s"," ","s","k","y"," ","t","h","e"]
示例 2：

输入：s = ["a"]
输出：["a"]
*/

// 这需要记住一个规律
// 1.翻转整个字节数组
// 2.逐个单词内部翻转
func reverseWords2(s []byte) {
	n := len(s)
	if n == 0 {
		return
	}
	reverseOne(s, 0, n-1)
	start := 0
	for i := 0; i < n; i++ {
		if s[i] == ' ' {
			reverseOne(s, start, i-1)
			start = i + 1
		} else if i == n-1 {
			reverseOne(s, start, i)
		}
	}
}

func reverseOne(s []byte, start, end int) {
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}
