package hash

/**
leetcode: 383. 赎金信
给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。

如果可以，返回 true ；否则返回 false 。

magazine 中的每个字符只能在 ransomNote 中使用一次。
*/

// hash key 是字符，value 是出现的次数
func canConstruct(ransomNote string, magazine string) bool {
    m := make(map[rune]int)
    for _, char := range magazine {
        m[char] += 1
    }
    for _, char := range ransomNote {
        if m[char] > 0 {
            m[char] -= 1
        } else {
            return false
        }
    }
    return true
}
