package math

/**
69. x 的平方根
给你一个非负整数 x ，计算并返回 x 的 算术平方根 。

由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。

注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
*/

// 二分法，想象一个数组，从 0 到 x 顺序递增，其中有一个一定是答案，用二分法寻找就可以了
func mySqrt(x int) int {
    var result = -1
    l, r := 0, x
    for r >= l {
        mid := (r-l)/2 + l
        if mid*mid <= x {
            result = mid
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    return result
}
