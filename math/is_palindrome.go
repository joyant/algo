package math

/**
9. 回文数
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文数
是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

例如，121 是回文，而 123 不是。
*/

// 1. 用字符串解决；先转为字符串，然后计算 palindrome 的值，比较两个字符串是否想等，或者两个指针指向字符串头尾，然后移动比较
// 2. 数学方法，每次取原数字的最后一位，构建对应的 palindrome，最后比较
// 数学方法可以用 11和 121 作为例子推演
func isPalindrome(x int) bool {
    if x < 0 || (x%10 == 0 && x != 0) {
        return false
    }
    reverse := 0
    for x > reverse {
        reverse = reverse*10 + x%10
        x /= 10
    }
    return x == reverse || x == reverse/10
}
