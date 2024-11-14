package array_and_string

func isPalindrome(s string) bool {
    n := len(s)
    i, j := 0, n-1
    for i < j {
        for ; i < n; i++ {
            if isLetter(s[i]) {
                break
            }
        }
        for ; j >= 0; j-- {
            if isLetter(s[j]) {
                break
            }
        }
        x, y := toUpper(s[i]), toUpper(s[j])
        if i >= n || j < 0 || x != y {
            return false
        }
        i++
        j--
    }
    return true
}

func isLetter(a byte) bool {
    return (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z')
}

func toUpper(a byte) byte {
    if a >= 'a' && a <= 'z' {
        return 'A' + (a - 'a')
    }
    return a
}
