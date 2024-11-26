package array_and_string

import "testing"

func Test_stringShift(t *testing.T) {
    x := stringShift("abc", [][]int{{1, 1}})
    t.Log(x)
    y := stringShift("abc", [][]int{{0, 1}})
    t.Log(y)
}
