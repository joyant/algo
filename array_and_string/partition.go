package array_and_string

/**
给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是
回文串。返回 s 所有可能的分割方案。

示例 1：

输入：s = "aab"
输出：[["a","a","b"],["aa","b"]]
示例 2：

输入：s = "a"
输出：[["a"]]

提示：

1 <= s.length <= 16
s 仅由小写英文字母组成
*/

func partition(s string) [][]string {
    res := make([][]string, 0)
    backtrack(s, []string{}, &res)
    return res
}

func isPalindrome2(s string) bool {
    for i := 0; i < len(s)/2; i++ {
        if s[i] != s[len(s)-i-1] {
            return false
        }
    }
    return true
}

func backtrack(s string, path []string, res *[][]string) {
    if len(s) == 0 {
        tmp := make([]string, len(path))
        copy(tmp, path)
        *res = append(*res, tmp)
        return
    }
    for i := 1; i <= len(s); i++ {
        if isPalindrome2(s[:i]) {
            path = append(path, s[:i])
            backtrack(s[i:], path, res)
            path = path[:len(path)-1]
        }
    }
}
