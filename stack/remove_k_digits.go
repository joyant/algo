package stack

import "strings"

/**
402. 移掉 K 位数字
level: medium
给你一个以字符串表示的非负整数 num 和一个整数 k ，移除这个数中的 k 位数字，使得剩下的数字最小。请你以字符串形式返回这个最小的数字。


示例 1 ：

输入：num = "1432219", k = 3
输出："1219"
解释：移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219 。
示例 2 ：

输入：num = "10200", k = 1
输出："200"
解释：移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。
示例 3 ：

输入：num = "10", k = 2
输出："0"
解释：从原数字移除所有的数字，剩余为空就是 0 。
*/

/**
我们需要从数字字符串中移除 k 个数字，使得剩下的数字最小，这就需要我们尽可能地保留低位的小数字。因此，我们可以使用一个单调递增的栈来解决这个问题。

具体的解题思路如下：

创建一个空栈，用于存放最终的数字序列，同时设置一个变量 removed 用于记录已经移除的数字的数量。

遍历数字字符串，对于每一个数字，如果该数字小于栈顶元素，并且 removed 小于 k，我们就移除栈顶元素，同时 removed 加 1。然后将当前数字入栈。这样可以保证栈内的数字序列是单调递增的。

如果遍历完数字字符串后，removed 仍小于 k，这说明我们还需要移除一些数字。由于此时栈内的数字序列已经是单调递增的，因此我们只需要从栈顶开始，继续移除剩余的数字。

最后，我们将栈内的数字序列转化为字符串，同时去除可能存在的前导零。

总结：
1. for 循环不断的把遍历到的最小的值往前移动，能移动多少移动多少（k），把更小的值往前移动，会使得结果更小
（比如1这个数字在百位上肯定比在各位上威力更大--让整个数字更小）
2. num 中的字符是无序的，所以未必 for 循环一定会执行，或者一定会将 k 递减至0，因此一定要加上这一句：
stack = stack[:len(stack)-k]
3. 去掉前导0
*/
func removeKDigits(num string, k int) string {
    stack := []byte{}
    for i := range num {
        digit := num[i]
        for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
            stack = stack[:len(stack)-1]
            k--
        }
        stack = append(stack, digit)
    }

    // remove the remaining digits from the tail.
    stack = stack[:len(stack)-k]
    // remove the leading zeros.
    ans := strings.TrimLeft(string(stack), "0")

    if len(ans) == 0 {
        return "0"
    }
    return ans
}
