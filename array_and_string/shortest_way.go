package array_and_string

/**
1055. 形成字符串的最短路径

对于任何字符串，我们可以通过删除其中一些字符（也可能不删除）来构造该字符串的 子序列 。(例如，“ace” 是 “abcde” 的子序列，而 “aec” 不是)。

给定源字符串 source 和目标字符串 target，返回 源字符串 source 中能通过串联形成目标字符串 target 的 子序列 的最小数量 。如果无法通过串联源字符串中的子序列来构造目标字符串，则返回 -1。



示例 1：

输入：source = "abc", target = "abcbc"
输出：2
解释：目标字符串 "abcbc" 可以由 "abc" 和 "bc" 形成，它们都是源字符串 "abc" 的子序列。
示例 2：

输入：source = "abc", target = "acdbc"
输出：-1
解释：由于目标字符串中包含字符 "d"，所以无法由源字符串的子序列构建目标字符串。
示例 3：

输入：source = "xyz", target = "xzyxz"
输出：3
解释：目标字符串可以按如下方式构建： "xz" + "y" + "xz"。

*/

/**
这个问题可以使用贪心算法来解决。我们可以遍历目标字符串target，每次都在源字符串source中找到对应的字符，然后移动到下一个字符。如果在源字符串中找不到对应的字符，那么就需要重新从源字符串的开始位置查找，这样相当于我们又使用了一次源字符串。如果目标字符串中有源字符串中不存在的字符，那么返回-1。
*/
// 1. 读懂题很重要
// 2. 两个指针，一个指向 source，一个指向 target，指向 source 的指针可以循环指向开头，只是每次指向开头都要判断一些东西
func shortestWay(source string, target string) int {
    sourceLen, targetLen := len(source), len(target)
    sourceIndex, targetIndex, res := 0, 0, 1
    atLeastOne := false

    for targetIndex < targetLen {
        if sourceIndex >= sourceLen {
            if !atLeastOne {
                return -1
            }
            atLeastOne = false
            sourceIndex = 0
            res++
        }
        if source[sourceIndex] == target[targetIndex] {
            atLeastOne = true
            sourceIndex++
            targetIndex++
        } else {
            sourceIndex++
        }
    }

    return res
}
