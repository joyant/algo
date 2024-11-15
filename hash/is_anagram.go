package hash

/**
leetcode: 242. 有效的字母异位词

给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false


提示:

1 <= s.length, t.length <= 5 * 104
s 和 t 仅包含小写字母
*/

/**
这个问题的关键是理解什么是字母异位词。字母异位词指的是两个字符串，它们的字母和数量都相同，但是顺序可能不同。例如，"anagram" 和 "nagaram" 就是字母异位词，因为它们都包含相同数量的 'a', 'n', 'g', 'r', 'm'。
知道了题干的要求，再提示用 map 解决问题，就可以了

至于：
进阶: 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

当前算法也是没有问题的，因为 rune 的本质是int32，已经可以支持unicode 字符了
*/
func isAnagram(s string, t string) bool {
    m := make(map[rune]int)
    for _, c := range s {
        m[c]++
    }
    for _, c := range t {
        m[c]--
    }
    for _, v := range m {
        if v != 0 {
            return false
        }
    }
    return true
}
