package binary

import "fmt"

/**
67. 二进制求和
给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和。
*/

// 从后往前遍历，逐个字符以二进制值相加，记录 carry
func addBinary(a string, b string) string {
    m, n := len(a), len(b)
    i, j := m-1, n-1
    var carry byte
    var result string
    for i >= 0 || j >= 0 || carry > 0 {
        var x, y byte
        if i >= 0 {
            x = a[i] - '0'
            i--
        }
        if j >= 0 {
            y = b[j] - '0'
            j--
        }
        sum := x + y + carry
        carry = sum / 2
        result = fmt.Sprintf("%d%s", sum%2, result)
    }
    return result
}
