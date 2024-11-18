package math

/**
66. 加一
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。
*/

// 不知道是不是我没看懂题，我的答案能通过，但和官方不一样，先这样吧，以后再看看
// todo
func plusOne(digits []int) []int {
    n := len(digits)
    if n == 0 {
        return []int{}
    }
    arr := make([]int, n+1)
    copy(arr[1:], digits)
    var carry = 0
    for i := n; i >= 0; i-- {
        add := 0
        if i == n {
            add = 1
        }
        sum := arr[i] + add + carry
        carry = sum / 10
        arr[i] = sum % 10
    }
    if arr[0] == 0 {
        return arr[1:]
    }
    return arr
}
