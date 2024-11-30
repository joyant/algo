package array_and_string

import "testing"

func Test_validWordSquare(t *testing.T) {
    x := validWordSquare([]string{"ball", "asee", "let", "lep"})
    t.Log(x)
    y := validWordSquare([]string{"abcd", "bnrt", "crmy", "dtye"})
    t.Log(y)
}
