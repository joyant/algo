package array_and_string

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func Test_reverseWord(t *testing.T) {
    x := reverseWords("the sky is blue")
    assert.Equal(t, "blue is sky the", x)
}
