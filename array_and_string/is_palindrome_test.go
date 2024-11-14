package array_and_string

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func Test_isPalindrome(t *testing.T) {
    x := isPalindrome("A man, a plan, a canal: Panama")
    assert.True(t, x)
}
