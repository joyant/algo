package hash

import "strings"

/**
leetcode: 290. 单词规律

给定一种规律 pattern 和一个字符串 s ，判断 s 是否遵循相同的规律。

这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 s 中的每个非空单词之间存在着双向连接的对应规律。

示例1:

输入: pattern = "abba", s = "dog cat cat dog"
输出: true
示例 2:

输入:pattern = "abba", s = "dog cat cat fish"
输出: false
示例 3:

输入: pattern = "aaaa", s = "dog cat cat dog"
输出: false
*/

// 这个题和 205 相似，都是用两个 map 来判断
func wordPattern(pattern string, s string) bool {
    m := make(map[string]byte)
    n := make(map[byte]string)
    slice := strings.Split(s, " ")
    if len(pattern) != len(slice) {
        return false
    }
    for i := 0; i < len(pattern); i++ {
        x, y := pattern[i], slice[i]
        if (m[y] > 0 && m[y] != x) || (n[x] != "" && n[x] != y) {
            return false
        }
        m[y] = x
        n[x] = y
    }
    return true
}
