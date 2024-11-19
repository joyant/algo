package sort

import "testing"

func Test_simpleBucket(t *testing.T) {
    arr := []int{1, 3, 4, 5}
    x := simpleBucket(arr)
    t.Log(x)
}
