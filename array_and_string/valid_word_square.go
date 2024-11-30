package array_and_string

func validWordSquare(words []string) bool {
    n := len(words)
    if n == 0 {
        return false
    }
    for i := 0; i < n; i++ {
        cur := words[i]
        for j := i; j < n; j++ {
            if i >= len(words[j]) {
                return true
            }
            a := cur[j]
            b := words[j][i]
            if a != b {
                return false
            }
        }
    }
    return true
}
