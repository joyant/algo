package array_and_string

func lengthOfLastWord(s string) int {
    // 思路是记录下连续不为空的字符串的长度
    var result int
    var length int
    for _, e := range s {
        if e == ' ' {
            result = max(result, length)
            length = 0
        } else {
            length++
        }
    }
    return max(result, length)
}
