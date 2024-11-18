package binary

/**
191. 位1的个数
给定一个正整数 n，编写一个函数，获取一个正整数的二进制形式并返回其二进制表达式中设置位的个数（也被称为汉明重量）。
*/

func hammingWeight(n int) int {
    // golang 环境中，不知道 n 是 32 还是 64位，如果知道就好了，这样就不需要修改 n 的值
    // 但不知道多少 bit 也没有关系，使用修改它的方法就可以
    sum := 0
    for n != 0 {
        sum += (n & 1)
        n >>= 1
    }
    return sum
}
