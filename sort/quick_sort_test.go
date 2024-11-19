package sort

import (
    "testing"
)

func Test_quickSort(t *testing.T) {
    arr := []int{10, 7, 8, 9, 1, 5}
    quickSort(arr, 0, len(arr)-1)
    t.Log(arr)
}

func Test_quickSort2(t *testing.T) {
    arr := []int{30, 20, 40}
    quickSort2(arr, 0, len(arr)-1)
    t.Log(arr)
}
