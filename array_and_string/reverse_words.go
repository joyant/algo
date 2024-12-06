package array_and_string

/**
leetcode: 151 反转字符串中的单词
给你一个字符串 s ，请你反转字符串中 单词 的顺序。

单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。

返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。

注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。

示例 1：
输入：s = "the sky is blue"
输出："blue is sky the"

示例 2：
输入：s = "  hello world  "
输出："world hello"
解释：反转后的字符串中不能存在前导空格和尾随空格。

示例 3：
输入：s = "a good   example"
输出："example good a"
解释：如果两个单词间有多余的空格，反转后的字符串需要将单词间的空格减少到仅有一个。
*/

// 从后往前遍历，记录第一个不为空格的 index，遇到空格，判断是否已经越过了一个完整的单词
func reverseWords(s string) string {
	n := len(s)
	var result []byte
	// index是一个完整单词的开始位置，比如它是 good 的 d 的位置
	var index = -1

	for i := n - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if index > -1 {
				if len(result) > 0 {
					result = append(result, ' ')
				}
				result = append(result, s[i+1:index+1]...)
			}
			index = -1
		} else if index == -1 {
			index = i
		}
	}
	if s[0] != ' ' && index > -1 {
		if len(result) > 0 {
			result = append(result, ' ')
		}
		result = append(result, s[:index+1]...)
	}
	return string(result)
}
