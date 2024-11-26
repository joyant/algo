package array_and_string

/**
1427. 字符串的左右移
给定一个包含小写英文字母的字符串 s 以及一个矩阵 shift，其中 shift[i] = [direction, amount]：

direction 可以为 0 （表示左移）或 1 （表示右移）。
amount 表示 s 左右移的位数。
左移 1 位表示移除 s 的第一个字符，并将该字符插入到 s 的结尾。
类似地，右移 1 位表示移除 s 的最后一个字符，并将该字符插入到 s 的开头。
对这个字符串进行所有操作后，返回最终结果。
*/

// 这里写的是容易 debug 的代码，实际做题的时候，可以少写一点变量
func stringShift(s string, shift [][]int) string {
    n := len(s)
    if n == 0 {
        return ""
    }
    move := 0
    for _, s := range shift {
        if s[0] == 0 {
            move -= s[1]
        } else {
            move += s[1]
        }
    }
    move = move % n
    if move == 0 {
        return s
    }
    if move < 0 {
        x := -move
        a, b := s[:x], s[x:]
        return b + a
    } else {
        a, b := s[n-move:], s[:n-move]
        return a + b
    }
}
