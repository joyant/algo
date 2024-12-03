package bit

/**
342. 4的幂
level: easy

给定一个整数，写一个函数来判断它是否是 4 的幂次方。如果是，返回 true ；否则，返回 false 。

整数 n 是 4 的幂次方需满足：存在整数 x 使得 n == 4x



示例 1：

输入：n = 16
输出：true
示例 2：

输入：n = 5
输出：false
示例 3：

输入：n = 1
输出：true
*/

// 三个条件
// 1. 大于0
// 2. 是2的幂，这就是一个公式，一个工具：n&(n-1) == 0
// 3. 特征：n%3 == 1
func isPowerOfFour(n int) bool {
    return n > 0 && n&(n-1) == 0 && n%3 == 1
}
