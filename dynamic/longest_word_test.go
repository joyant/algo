package dynamic

import "testing"

func Test_longestWord(t *testing.T) {
    words := []string{"cat", "banana", "dog", "nana", "walk", "walker", "dogwalker"}
    x := longestWord(words)
    t.Logf("x:%s", x)
}
