package hash

/**
leetcode: 205. 同构字符串

给定两个字符串 s 和 t ，判断它们是否是同构的。

如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。

每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。

示例 1:

输入：s = "egg", t = "add"
输出：true
示例 2：

输入：s = "foo", t = "bar"
输出：false
示例 3：

输入：s = "paper", t = "title"
输出：true

*/
// 用两个 map 解决，一个用来存映射关系，另一个用来验证
func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    m, n := make(map[byte]byte), make(map[byte]byte)
    for i := 0; i < len(s); i++ {
        x, y := s[i], t[i]
        if (m[x] > 0 && m[x] != y) || (n[y] > 0 && n[y] != x) {
            return false
        }
        m[x] = y
        n[y] = x
    }
    return true
}
