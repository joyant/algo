package slow_fast_pointer

/**
leetcode: 202. 快乐数

编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」 定义为：

对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
如果这个过程 结果为 1，那么这个数就是快乐数。
如果 n 是 快乐数 就返回 true ；不是，则返回 false 。



示例 1：

输入：n = 19
输出：true
解释：
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1
示例 2：

输入：n = 2
输出：false
*/

/**
我觉得这个题应该是中等，甚至是困难
快慢指针的一个规律，一开始慢指针不动，快指针走一步，然后每次循环，慢指针走一步，快指针走两步
*/
func isHappy(n int) bool {
    slow, fast := n, next(n)
    for fast != 1 && slow != fast {
        slow = next(slow)
        fast = next(next(fast))
    }
    return fast == 1
}

func next(n int) int {
    sum := 0
    for n > 0 {
        digit := n % 10
        sum += digit * digit
        n /= 10
    }
    return sum
}
