package dynamic

import "testing"

func Test_longestPalindrome(t *testing.T) {
	x := longestPalindrome("babad")
	t.Log(x)
	y := longestPalindrome("cbbd")
	t.Log(y)
}
