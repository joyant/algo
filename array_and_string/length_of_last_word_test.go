package array_and_string

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestLengthOfLastWord(t *testing.T) {
    x := lengthOfLastWord("Today is a nice day")
    assert.Equal(t, x, 6)
}
