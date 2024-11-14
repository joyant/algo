package array_and_string

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func Test_strStr(t *testing.T) {
    x := strStr("sadbutsad", "sad")
    assert.Equal(t, 0, x)
    y := strStr("leetcode", "leeto")
    assert.Equal(t, -1, y)
}
