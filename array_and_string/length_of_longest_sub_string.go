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

/**
下面是一个例子，说明：lastSeenPos >= start 判断不是多余的

假设我们有一个字符串 "abba"，我们来逐步执行这个函数：

1. 初始化 `start = 0, maxLen = 0`，并创建一个空的哈希表 `m`。

2. 首先处理字符 'a'，哈希表中没有 'a'，所以将 'a' 和其位置 0 加入哈希表，`maxLen` 更新为 `1`。

3. 然后处理字符 'b'，哈希表中没有 'b'，所以将 'b' 和其位置 1 加入哈希表，`maxLen` 更新为 `2`。

4. 接下来再次处理字符 'b'，哈希表中有 'b'，并且其最后出现的位置 `1 >= start`，所以移动 `start` 到 `2`，并更新 'b' 在哈希表中的位置为 `2`，`maxLen` 保持不变。

5. 最后处理字符 'a'，哈希表中有 'a'，但是其最后出现的位置 `0 < start`，所以并不需要移动 `start`。然后更新 'a' 在哈希表中的位置为 `3`，`maxLen` 保持不变。

所以最长的无重复字符的子串长度为 `2`。

从这个例子中，你可以看到 `lastSeenPos >= start` 这个判断的作用。在处理第二个 'b' 和最后一个 'a' 时，这个判断都起到了作用。对于第二个 'b'，
由于它在当前滑动窗口内已经出现过，所以需要移动 `start`。而对于最后一个 'a'，虽然它之前出现过，但是并没有在当前滑动窗口内出现过，所以并不需要移动 `start`。
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
