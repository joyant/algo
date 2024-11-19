package sort

import "testing"

func Test_countingSortNaive(t *testing.T) {
    nums := []int{1, 5, 4, 3}
    countingSortNaive(nums)
    t.Log(nums)
}
